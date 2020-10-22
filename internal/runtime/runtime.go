package runtime

import (
	"context"
	"sync"
	"time"

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
	conf *Config

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
	done   bool
	logger zerolog.Logger
	sync.Mutex
}

// Start the empirica recruitment runtime
func Start(conf *Config, conn *storage.Conn, mturk, mturkSandbox *mturk.Session) (*Runtime, error) {

	logger := log.With().Str("pkg", "runtime").Logger()

	if !conf.Debug {
		logger = logger.Level(zerolog.Disabled)
	}

	r := &Runtime{
		conf:         conf,
		conn:         conn,
		mturk:        mturk,
		mturkSandbox: mturkSandbox,
		runs:         make(map[string]*runState),
		updates:      make(chan string),
		triggers:     make(chan *runState),
		logger:       logger,
	}

	if !conf.Disable {
		go r.processRuns()
		// go r.processEvents()
		err := r.registerExistingSteps()
		if err != nil {
			return nil, err
		}
		r.registerHooks()

		r.logger.Debug().Msg("Runtime started")
	}

	return r, nil
}

// Stop the empirica recruitment runtime
func (r *Runtime) Stop() {
	r.done = true
	r.logger.Debug().Msg("Runtime stopped")
	// stop run loop here?
}

func (r *Runtime) addRun(run *ent.Run) {
	_, ok := r.runs[run.ID]
	if ok {
		// r.logger.Warn().Msgf("Run %s already tracked", run.ID)
		return
	}

	// If the status is created, and scheduled, set trigger for startAt
	if run.Status == runModel.StatusCREATED {
		if run.StartAt != nil {
			r.logger.Debug().Str("runID", run.ID).Msg("Adding scheduled run")
			state, err := r.newRunState(run)
			if err != nil {
				r.logger.Error().Err(err).Str("runID", run.ID).Msg("Failed to add scheduled run")
			}
			state.scheduleNext()
		}
		return
	}

	// Double-check this can only happen on created or running Runs
	if run.Status != runModel.StatusRUNNING {
		return
	}

	r.logger.Debug().Str("runID", run.ID).Msg("Adding unscheduled run")
	state, err := r.newRunState(run)
	if err != nil {
		r.logger.Error().Err(err).Str("runID", run.ID).Msg("Failed to add unscheduled run")
	}
	state.scheduleNext()
}

func (r *Runtime) removeRun(run *ent.Run) {
	r.logger.Debug().Msg("Removing run")

	state, ok := r.runs[run.ID]
	if !ok {
		return
	}

	delete(r.runs, run.ID)

	if state.timer != nil && !state.timer.Stop() {
		<-state.timer.C
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
	case runModel.StatusCREATED, runModel.StatusRUNNING:
		r.addRun(runRecord)
	case runModel.StatusDONE, runModel.StatusFAILED, runModel.StatusTERMINATED, runModel.StatusPAUSED:
		r.removeRun(runRecord)
	default:
		r.logger.Error().Msgf("unknown run status: %s", runRecord.Status.String())
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
					go func() {
						ID, exists := m.ID()
						if exists {
							r.updates <- ID
						}
					}()
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

	for _, runID := range runIDs {
		r.updates <- runID
	}

	return nil
}
