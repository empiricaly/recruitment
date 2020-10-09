package runtime

import (
	"context"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	steprunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (r *Runtime) startRun(run *ent.Run) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	template, err := run.QueryTemplate().Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("start run: query template")
		return
	}

	steps, err := template.QuerySteps().Order(ent.Asc(stepModel.FieldIndex)).All(ctx)
	if err != nil {
		log.Error().Err(err).Msg("start run: query steps")
		return
	}

	tx, err := r.conn.Tx(ctx)
	if err != nil {
		log.Error().Err(err).Msg("start run: create transaction")
		return
	}

	for _, step := range steps {
		_, err := tx.
			StepRun.
			Create().
			SetID(xid.New().String()).
			SetRun(run).
			SetStep(step).
			SetStatus(steprunModel.StatusCREATED).
			SetParticipantsCount(0).
			Save(ctx)
		if err != nil {
			err = errors.Wrap(err, "start run: create stepRun")
			return
		}
	}

	run, err = tx.Run.UpdateOne(run).
		SetStatus(runModel.StatusRUNNING).
		SetStartAt(time.Now()).
		Save(ctx)

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("start run: commit transaction")
		return
	}

	r.startStep(run)
}

func (r *Runtime) endRun(run *ent.Run) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	tx, err := r.conn.Tx(ctx)
	if err != nil {
		log.Error().Err(err).Msg("end run: create transaction")
		return
	}

	now := time.Now()

	currentRun, err := run.QueryCurrentStep().Only(ctx)
	if err != nil {
		log.Error().Err(err).Msg("end run: get last step")
		return
	}

	_, err = tx.
		StepRun.
		UpdateOne(currentRun).
		SetStatus(steprunModel.StatusCREATED).
		SetEndedAt(now).
		Save(ctx)
	if err != nil {
		log.Error().Err(err).Msg("end run: update last step")
		return
	}

	_, err = tx.
		Run.
		UpdateOne(run).
		SetStatus(runModel.StatusRUNNING).
		SetStartAt(now).
		Save(ctx)
	if err != nil {
		log.Error().Err(err).Msg("end run: update run")
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msg("end run: commit transaction")
		return
	}
}

func (r *Runtime) endStep(run *ent.Run, stepRun *ent.StepRun) {
	// TODO end anything MTurk is doing for this run.
	// On a HIT Step, we should make sure the HIT is deleted and people who
	// haven't  completed the HIT yet are not brought to the next round
}

func (r *Runtime) startStep(run *ent.Run) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	currentStepRun, err := run.QueryCurrentStep().WithStep().Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error().Err(err).Msg("start step: get current step")
		return
	}

	stepRuns, err := run.QuerySteps().WithStep().All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error().Err(err).Msg("start step: get current step")
		return
	}

	if len(stepRuns) == 0 {
		log.Error().Msgf("trying to run a Run without Steps! (%s)", run.ID)
		return
	}

	if currentStepRun != nil {
		r.endStep(run, currentStepRun)
	}

	if err := ent.WithTx(ctx, r.conn.Client, func(tx *ent.Tx) error {
		var nextStepRun *ent.StepRun
		if currentStepRun == nil {
			// starting the first run

			for _, stepRun := range stepRuns {
				if stepRun.Edges.Step.Index == 0 {
					nextStepRun = stepRun
					break
				}
			}

			if nextStepRun == nil {
				return errors.Errorf("could not find first step for run (%s)", run.ID)
			}
		} else {
			// running subsequent run

			for _, stepRun := range stepRuns {
				if stepRun.Edges.Step.Index == currentStepRun.Edges.Step.Index {
					nextStepRun = stepRun
					break
				}
			}

			if nextStepRun == nil {
				return errors.Errorf("could not find next step for run (run: %s, prev step: %s)", run.ID, currentStepRun.ID)
			}
		}

		now := time.Now()

		_, err = tx.
			StepRun.
			UpdateOne(nextStepRun).
			SetStatus(steprunModel.StatusCREATED).
			SetEndedAt(now).
			Save(ctx)
		if err != nil {
			return errors.Wrapf(err, "update last step", run.ID)
		}

		err = r.mturk.RunStep(run, nextStepRun)
		if err != nil {
			return errors.Wrap(err, "run mturk for step")
		}

		return nil
	}); err != nil {
		log.Error().Err(err).Msg("start step: transaction failed")
	}
}
