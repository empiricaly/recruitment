package runtime

import (
	"context"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/hook"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Runtime manages the empirica recruitment run loop that will trigger timed
// events as needed (start run, go to next step, etc.)
type Runtime struct {
	// db conn
	conn *storage.Conn

	// runs tracked by the runtime. The map key is the Run ID.
	runs map[string]*runState

	// update wants Run IDs to update the runtime
	// When a new Run ID comes in, it is fetched from the DB, then
	// if it is already tracked, we check what change that implies
	// for the runtime, or if it doesn't exist, we add it to the runtime.
	updates chan string

	// triggers receives timed events, at the time of the event, for the event
	// to be processed.
	triggers chan *runState

	// If done == true, stop running hooks
	done bool
}

// Start the empirica recruitment runtime
func Start(conn *storage.Conn) (*Runtime, error) {
	r := &Runtime{
		conn:     conn,
		runs:     make(map[string]*runState),
		updates:  make(chan string),
		triggers: make(chan *runState),
	}
	go r.processRuns()
	go r.processEvents()
	err := r.registerExistingSteps()
	if err != nil {
		return nil, err
	}
	r.registerHooks()

	return r, nil
}

// Stop the empirica recruitment runtime
func (r *Runtime) Stop() {
	r.done = true
	// stop run loop here?
}

//go:generate stringer  -linecomment -type=runEvent  -output=runtime_strings.go

type runEvent int

const (
	startRunEvent runEvent = iota // Start Run
	endRunEvent                   // End Run
	nextStepEvent                 // Next Run Step
)

type runState struct {
	run       *ent.Run
	timer     *time.Timer
	nextEvent runEvent
}

func (r *Runtime) addRun(run *ent.Run) {
	return
	_, ok := r.runs[run.ID]
	if ok {
		log.Warn().Msgf("run %s already tracked", run.ID)
		return
	}

	// If the status is created, and scheduled, set trigger for startAt
	if run.Status == runModel.StatusCREATED {
		if run.StartAt != nil {
			r.schedule(run, startRunEvent, time.Until(*run.StartAt))
		}
		return
	}

	// Double-check this can only happen on created or running Runs
	if run.Status != runModel.StatusRUNNING {
		log.Warn().Msgf("runtime: tried to add run with wrong status: %s", run.ID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	currentStepRun, err := run.QueryCurrentStep().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error().Err(err).Msg("query currentStepRun for next event trigger")
		return
	}

	// Run set to start now
	if currentStepRun == nil {
		if run.StartAt == nil {
			run.Start()
		} else {
			log.Error().Msg("Run is running but first step hasn't started yet and is scheduled")
		}
		return
	}

	r.scheduleNextStep(run)
}

func (r *Runtime) scheduleNextStep(run *ent.Run) {
	// Get related objects
	_, currentStep, _, _, steps, err := run.Relations()
	if err != nil {
		log.Error().Err(err).Msg("preparing next event trigger")
		return
	}

	var next bool
	var nextStep *ent.Step
	for _, step := range steps {
		if next {
			nextStep = step
		}
		if step.ID == currentStep.ID {
			next = true
		}
	}

	// If there is no next step, it's the last step
	if nextStep == nil {
		until, err := untilNextStep(run, steps, steps[len(steps)-1], true)
		if err != nil {
			log.Error().Err(err).Msg("getting end of run time")
		}
		r.schedule(run, endRunEvent, until)
		return
	}

	// Schedule next step
	until, err := untilNextStep(run, steps, steps[len(steps)-1], false)
	if err != nil {
		log.Error().Err(err).Msg("getting next step start")
	}
	r.schedule(run, nextStepEvent, until)
}

func (r *Runtime) schedule(run *ent.Run, nextEvent runEvent, wait time.Duration) {
	_, ok := r.runs[run.ID]
	if ok {
		log.Warn().Msgf("run %s already scheduled", run.ID)
		return
	}
	// Schedule starting the Run
	state := &runState{
		run:       run,
		nextEvent: nextEvent,
	}

	state.timer = time.AfterFunc(wait, func() {
		r.triggers <- state
	})
	r.runs[run.ID] = state
}

func untilNextStep(run *ent.Run, steps []*ent.Step, step *ent.Step, toEnd bool) (time.Duration, error) {
	if run.StartedAt == nil {
		return 0, errors.New("run has not started")
	}
	var dur time.Duration
	for _, s := range steps {
		if !toEnd && s.ID == step.ID {
			break
		}
		dur += time.Second * time.Duration(s.Duration)
		if toEnd && s.ID == step.ID {
			break
		}
	}

	return time.Until(run.StartedAt.Add(dur)), nil
}

func (r *Runtime) removeRun(run *ent.Run) {
	state, ok := r.runs[run.ID]
	if !ok {
		return
	}

	if !state.timer.Stop() {
		<-state.timer.C
	}
}

func (r *Runtime) processEvent(state *runState) {
	switch state.nextEvent {
	case startRunEvent:
		state.run.Start()
	case endRunEvent:
		state.run.End()
	case nextStepEvent:
		state.run.RunStep()
		r.scheduleNextStep(state.run)
	default:
		log.Error().Msgf("connot process next event on %s: %s", state.run.ID, state.nextEvent.String())
	}
}

func (r *Runtime) processRun(runID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	runRecord, err := r.conn.Run.Get(ctx, runID)
	if err != nil {
		if ent.IsNotFound(err) {
			r.removeRun(runRecord)
		} else if ent.IsNotSingular(err) {
			log.Error().Err(err).Msg("multiple runs with same ID!")
		} else {
			log.Error().Err(err).Msg("updating Run state in runtime")
		}
		return
	}

	switch runRecord.Status {
	case runModel.StatusDONE, runModel.StatusFAILED, runModel.StatusTERMINATED, runModel.StatusPAUSED:
		r.removeRun(runRecord)
	case runModel.StatusCREATED, runModel.StatusRUNNING:
		r.addRun(runRecord)
	default:
		log.Error().Msgf("unknown run status: %s", runRecord.Status.String())
	}
}

func (r *Runtime) processEvents() {
	for {
		event := <-r.triggers
		r.processEvent(event)
	}
}

func (r *Runtime) processRuns() {
	for {
		runID := <-r.updates
		r.processRun(runID)
	}
}

func (r *Runtime) registerHooks() {
	r.conn.Run.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.RunFunc(func(ctx context.Context, m *ent.RunMutation) (ent.Value, error) {
				if r.done {
					return next.Mutate(ctx, m)
				}

				// After the update happened, go ahead and add Run to Runtime
				defer func() {
					ID, exists := m.ID()
					if exists {
						r.updates <- ID
					}
				}()

				return next.Mutate(ctx, m)
			})
		},
		ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne))
}

func (r *Runtime) registerExistingSteps() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	runIDs, err := r.conn.Run.Query().
		Where(runModel.StatusIn(runModel.StatusCREATED, runModel.StatusRUNNING)).
		Select(runModel.FieldID).
		Strings(ctx)
	if err != nil {
		return errors.Wrap(err, "initialize exisitng runs")
	}

	// log.Debug().Interface("runIDs", runIDs).Msg("got runs")

	for _, runID := range runIDs {
		r.updates <- runID
	}

	return nil
}
