package runtime

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/hook"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/mturk"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Runtime manages the empirica recruitment run loop that will trigger timed
// events as needed (start run, go to next step, etc.)
type Runtime struct {
	// db conn
	conn *storage.Conn

	// mturk session
	mturk        *mturk.Session
	mturkSandbox *mturk.Session

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

	logger zerolog.Logger
}

var initRand sync.Once

// Start the empirica recruitment runtime
func Start(conn *storage.Conn, mturk, mturkSandbox *mturk.Session) (*Runtime, error) {
	initRand.Do(func() {
		rand.Seed(time.Now().UnixNano())
	})

	r := &Runtime{
		conn:         conn,
		mturk:        mturk,
		mturkSandbox: mturkSandbox,
		runs:         make(map[string]*runState),
		updates:      make(chan string),
		triggers:     make(chan *runState),
		logger:       log.With().Str("pkg", "runtime").Logger(),
	}
	go r.processRuns()
	go r.processEvents()
	err := r.registerExistingSteps()
	if err != nil {
		return nil, err
	}
	r.registerHooks()

	r.logger.Debug().Msg("Runtime started")

	return r, nil
}

// Stop the empirica recruitment runtime
func (r *Runtime) Stop() {
	r.done = true
	r.logger.Debug().Msg("Runtime stopped")
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

	_, ok := r.runs[run.ID]
	if ok {
		// r.logger.Warn().Msgf("Run %s already tracked", run.ID)
		return
	}

	r.logger.Debug().Msg("Adding run")
	defer r.logger.Debug().Msg("Run added")

	// If the status is created, and scheduled, set trigger for startAt
	if run.Status == runModel.StatusCREATED {
		if run.StartAt != nil {
			r.logger.Debug().Msg("Scheduling run")
			r.schedule(run, startRunEvent, time.Until(*run.StartAt))
		}
		return
	}

	// Double-check this can only happen on created or running Runs
	if run.Status != runModel.StatusRUNNING {
		r.logger.Warn().Msgf("tried to add run with wrong status: %s", run.ID)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	currentStepRun, err := run.QueryCurrentStep().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		r.logger.Error().Err(err).Msg("Query currentStepRun for next event trigger")
		return
	}

	// Run set to start now
	if currentStepRun == nil {
		if run.StartAt == nil {
			// r.startRun(run)
			r.schedule(run, startRunEvent, time.Duration(0))
		} else {
			r.logger.Error().Time("startAt", *run.StartAt).Msg("Run is running but first step hasn't started yet and is scheduled")
		}
		return
	}

	r.scheduleNextStep(run)
}

func (r *Runtime) scheduleNextStep(run *ent.Run) {
	r.logger.Debug().Msg("Scheduling next step")

	// Get related objects
	_, currentStep, _, _, steps, err := run.Relations()
	if err != nil {
		r.logger.Error().Err(err).Msg("Preparing next event trigger")
		return
	}

	spew.Dump(currentStep, steps)

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
		r.logger.Debug().Msg("Scheduling end of run")

		until, err := untilNextStep(run, steps, steps[len(steps)-1], true)
		if err != nil {
			r.logger.Error().Err(err).Msg("Getting end of run time")
			return
		}
		r.schedule(run, endRunEvent, until)
		return
	}

	// Schedule next step
	until, err := untilNextStep(run, steps, steps[len(steps)-1], false)
	if err != nil {
		r.logger.Error().Err(err).Msg("getting next step start")
		return
	}

	r.schedule(run, nextStepEvent, until)
	r.logger.Debug().Msg("Next step scheduled")
}

func (r *Runtime) schedule(run *ent.Run, nextEvent runEvent, wait time.Duration) {
	_, ok := r.runs[run.ID]
	if ok {
		r.logger.Warn().Msgf("run %s already scheduled", run.ID)
		return
	}
	// Schedule starting the Run
	state := &runState{
		run:       run,
		nextEvent: nextEvent,
	}

	r.logger.Debug().
		Dur("wait", wait).
		Time("at", time.Now().Add(wait)).
		Interface("event", nextEvent.String()).
		Msg("scheduling")

	state.timer = time.AfterFunc(wait, func() {
		r.logger.Debug().Msg("Next step triggered")
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
		dur += time.Minute * time.Duration(s.Duration)
		if toEnd && s.ID == step.ID {
			break
		}
	}

	return time.Until(run.StartedAt.Add(dur)), nil
}

func (r *Runtime) removeRun(run *ent.Run) {
	r.logger.Debug().Msg("Removing run")

	state, ok := r.runs[run.ID]
	if !ok {
		return
	}

	delete(r.runs, run.ID)

	if !state.timer.Stop() {
		<-state.timer.C
	}
}

func (r *Runtime) processEvent(state *runState) {
	switch state.nextEvent {
	case startRunEvent:
		r.logger.Debug().Msg("Processing event: start run")
		r.startRun(state.run)
		run, err := r.conn.Run.Get(context.Background(), state.run.ID)
		if err == nil {
			state.run = run
		}
		r.scheduleNextStep(state.run)
	case endRunEvent:
		r.logger.Debug().Msg("Processing event: end run")
		r.endRun(state.run)
		r.removeRun(state.run)
	case nextStepEvent:
		r.logger.Debug().Msg("Processing event: next step")
		r.startStep(state.run)
		r.scheduleNextStep(state.run)
	default:
		r.logger.Error().Msgf("connot process next event on %s: %s", state.run.ID, state.nextEvent.String())
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
			r.logger.Error().Err(err).Msg("multiple runs with same ID!")
		} else {
			r.logger.Error().Err(err).Msg("updating Run state in runtime")
		}
		return
	}

	switch runRecord.Status {
	case runModel.StatusDONE, runModel.StatusFAILED, runModel.StatusTERMINATED, runModel.StatusPAUSED:
		r.removeRun(runRecord)
	case runModel.StatusCREATED, runModel.StatusRUNNING:
		r.addRun(runRecord)
	default:
		r.logger.Error().Msgf("unknown run status: %s", runRecord.Status.String())
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
				r.logger.Debug().Msg("hook triggered")
				if r.done {
					return next.Mutate(ctx, m)
				}

				// After the update happened, go ahead and add Run to Runtime
				go func() {
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

	// r.logger.Debug().Interface("runIDs", runIDs).Msg("got runs")

	for _, runID := range runIDs {
		r.updates <- runID
	}

	return nil
}
