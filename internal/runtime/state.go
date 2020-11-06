package runtime

import (
	"context"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/run"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	steprunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/pkg/errors"
)

//go:generate stringer  -linecomment -type=runEvent  -output=runtime_strings.go

type runEvent int

const (
	startRunEvent runEvent = iota // Start Run
	endRunEvent                   // End Run
	nextStepEvent                 // Next Step
)

type runState struct {
	timer     *time.Timer
	nextEvent runEvent

	run *ent.Run

	project        *ent.Project
	template       *ent.Template
	steps          []*ent.Step
	stepRuns       []*ent.StepRun
	currentStep    *ent.Step
	currentStepRun *ent.StepRun
	nextStep       *ent.Step
	nextStepRun    *ent.StepRun

	mturkSession *mturk.Session

	*Runtime
}

func (r *Runtime) newRunState(run *ent.Run) (*runState, error) {
	s, ok := r.runs[run.ID]
	if ok {
		return s, nil
	}

	// Schedule starting the Run
	state := &runState{
		Runtime: r,
		run:     run,
	}

	err := state.refresh()
	if err != nil {
		return nil, errors.Wrap(err, "could not add run state")
	}

	if len(state.steps) == 0 {
		return nil, errors.New("run has no steps")
	}

	state.setNextStep()

	if state.template.Sandbox {
		state.mturkSession = r.mturkSandbox
	} else {
		state.mturkSession = r.mturk
	}

	r.runs[run.ID] = state

	return state, nil
}

func (r *runState) scheduleNext() error {
	if r.timer != nil {
		return errors.New("run already scheduled")
	}

	var err error
	var nextEventType runEvent
	var nextEventWait time.Duration
	if r.currentStep == nil {
		nextEventType = startRunEvent
		if r.run.StartAt != nil {
			nextEventWait = time.Until(*r.run.StartAt)
		} else {
			if r.run.Status == run.StatusRUNNING {
				nextEventWait = time.Duration(0)
			} else {
				return errors.New("trying to schedule run that's not running and not scheduled")
			}
		}
	} else if r.nextStep == nil {
		if r.endOfRunReached() {
			nextEventType = endRunEvent
			nextEventWait, err = r.timeUntilEndOfStep()
			if err != nil {
				return errors.Wrap(err, "trying to schedule end of run but cannot find time to wait")
			}
		} else {
			return errors.New("no next step found")

		}
	} else {
		nextEventType = nextStepEvent
		nextEventWait, err = r.timeUntilEndOfStep()
		if err != nil {
			return errors.Wrap(err, "trying to schedule next step but cannot find time to wait")
		}
	}

	r.nextEvent = nextEventType

	r.logger.Debug().
		Str("waiting", nextEventWait.Round(time.Second).String()).
		Time("until", time.Now().Add(nextEventWait)).
		Interface("next event", nextEventType.String()).
		Msg("Scheduling next event")

	r.timer = time.AfterFunc(nextEventWait, r.processNextStep)

	return nil
}

func (r *runState) processNextStep() {
	currentEventType := r.nextEvent.String()
	r.logger.Debug().Msgf("Processing processing %s", currentEventType)
	if err := r.processStep(); err != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		r.refresh()

		at := time.Now()

		if r.currentStepRun != nil {
			_, err := r.currentStepRun.
				Update().
				SetStatus(steprunModel.StatusFAILED).
				SetEndedAt(at).
				Save(ctx)
			if err != nil {
				r.logger.Error().Err(err).Msg("Failed stopping failed step")
			}
		}

		_, err2 := r.run.Update().
			SetStatus(runModel.StatusFAILED).
			SetEndedAt(at).
			SetError(errors.Wrapf(err, "Failed processing %s", currentEventType).Error()).
			Save(ctx)
		if err2 != nil {
			r.logger.Error().Err(err2).Msg("Failed stopping failed run")
		}

		r.logger.Error().Err(err).Msgf("Failed processing %s", currentEventType)
		return
	}
	r.logger.Debug().Msgf("Finished processing %s", currentEventType)
}

func (r *runState) processStep() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	r.timer = nil

	var err error
	switch r.nextEvent {
	case startRunEvent:
		return errors.New("Oh nooooo")
		at := time.Now()
		err = r.startRun(ctx, at)
		if err != nil {
			return errors.Wrap(err, "start run")
		}
		err = r.advanceSteps(ctx, at)
		if err != nil {
			return errors.Wrap(err, "advance first step")
		}
		err = r.startStep(ctx, at)
		if err != nil {
			return errors.Wrap(err, "start first step")
		}
		err = r.scheduleNext()
		if err != nil {
			return errors.Wrap(err, "first schedule next")
		}
	case nextStepEvent:
		at := time.Now()
		err = r.endStep(ctx, at)
		if err != nil {
			return errors.Wrap(err, "end step")
		}
		err = r.advanceSteps(ctx, at)
		if err != nil {
			return errors.Wrap(err, "advance step")
		}
		err = r.startStep(ctx, at)
		if err != nil {
			return errors.Wrap(err, "start step")
		}
		err = r.scheduleNext()
		if err != nil {
			return errors.Wrap(err, "schedule next")
		}
	case endRunEvent:
		at := time.Now()
		err = r.endStep(ctx, at)
		if err != nil {
			return errors.Wrap(err, "end last step")
		}
		err = r.advanceSteps(ctx, at)
		if err != nil {
			return errors.Wrap(err, "advance step last")
		}
		err = r.endRun(ctx, at)
		if err != nil {
			return errors.Wrap(err, "end run")
		}
		// remove runState from Runtime? Already done through hook?
	default:
		return errors.Errorf("unknown next event: %v", r.nextEvent)
	}

	return nil
}

func (r *runState) endOfRunReached() bool {
	return r.currentStep.Index >= len(r.steps)-1
}

func (r *runState) setNextStep() error {
	if len(r.steps) == 0 {
		return errors.New("run has not steps")
	}

	if len(r.stepRuns) == 0 {
		return errors.New("run has not step runs")
	}

	if len(r.steps) != len(r.stepRuns) {
		return errors.New("step count != step runs count")
	}

	if r.currentStep == nil {
		r.nextStep = r.steps[0]
		r.nextStepRun = r.stepRuns[0]
		return nil
	}

	// Reached the end of the run
	if r.endOfRunReached() {
		r.nextStep = nil
		r.nextStepRun = nil
		return nil
	}

	r.nextStep = r.steps[r.currentStep.Index+1]
	r.nextStepRun = r.stepRuns[r.currentStep.Index+1]

	return nil
}

func (r *runState) refresh() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	var err error
	r.run, err = r.conn.Run.Get(ctx, r.run.ID)
	if err != nil {
		return errors.Wrap(err, "refresh: run")
	}

	if r.project == nil {
		r.project, err = r.run.QueryProject().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "refresh: project")
		}
	}

	if r.template == nil {
		r.template, err = r.run.QueryTemplate().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "refresh: template")
		}
	}

	if r.steps == nil {
		r.steps, err = r.template.QuerySteps().Order(ent.Asc(stepModel.FieldIndex)).All(ctx)
		if err != nil {
			return errors.Wrap(err, "refresh: steps")
		}
	}

	r.stepRuns, err = r.run.QuerySteps().Order(ent.Asc(stepModel.FieldIndex)).All(ctx)
	if err != nil {
		return errors.Wrap(err, "refresh: runsteps")
	}

	r.currentStepRun, err = r.run.QueryCurrentStep().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return errors.Wrap(err, "refresh: currentStepRun")
	}

	if r.currentStepRun == nil {
		return nil
	}

	r.currentStep, err = r.currentStepRun.QueryStep().Only(ctx)
	if err != nil {
		return errors.Wrap(err, "refresh: currentStep")
	}

	return nil
}

func (r *runState) timeUntilEndOfStep() (time.Duration, error) {
	if r.run.StartedAt == nil {
		return 0, errors.New("run has not started")
	}

	var dur time.Duration
	for _, s := range r.steps {
		dur += time.Minute * time.Duration(s.Duration)
		if s.ID == r.currentStep.ID {
			break
		}
	}

	return time.Until(r.run.StartedAt.Add(dur)), nil
}
