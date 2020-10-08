package runtime

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/hook"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Runtime manages the empirica recruitment run loop that will trigger timed
// events as needed (start run, go to next step, etc.)
type Runtime struct {
	// db conn
	conn *storage.Conn

	// runs tracked by the runtime. A map of Run ID to *Run.
	runs map[string]*ent.Run

	// update wants Run IDs to update the runtime
	// When a new Run ID comes in, it is fetched from the DB, then
	// if it is already tracked, we check what change that implies
	// for the runtime, or if it doesn't exist, we add it to the runtime.
	updates chan string

	// If done == true, stop running hooks
	done bool
}

// Start the empirica recruitment runtime
func Start(conn *storage.Conn) (*Runtime, error) {
	r := &Runtime{
		conn:    conn,
		runs:    make(map[string]*ent.Run),
		updates: make(chan string),
	}
	go r.processRuns()
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

func (r *Runtime) processRuns() {
	for {
		runID := <-r.updates
		fmt.Printf("\n\n\n\n\n\nUpdate run: %s\n\n\n\n", runID)
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
		Where(run.StatusIn(run.StatusCREATED, run.StatusRUNNING, run.StatusPAUSED)).
		Select(run.FieldID).
		Strings(ctx)
	if err != nil {
		return errors.Wrap(err, "initialize exisitng runs")
	}

	log.Debug().Interface("runIDs", runIDs).Msg("got runs")

	for _, runID := range runIDs {
		r.updates <- runID
	}

	return nil
}
