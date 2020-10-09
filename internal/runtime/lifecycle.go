package runtime

import (
	"context"
	"math"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	steprunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	templateModel "github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (r *Runtime) filterParticipants(ctx context.Context, tx *ent.Tx, template *ent.Template) ([]*ent.Participant, error) {
	participants, err := tx.Participant.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "find participants")
	}

	// TODO should filter participants here
	l := math.Min(float64(template.ParticipantCount), float64(len(participants)))
	return participants[:int(l)], nil
}

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

	if err := ent.WithTx(ctx, r.conn.Client, func(tx *ent.Tx) error {

		for i, step := range steps {
			stepRun, err := tx.
				StepRun.
				Create().
				SetID(xid.New().String()).
				SetRun(run).
				SetStep(step).
				SetStatus(steprunModel.StatusCREATED).
				SetParticipantsCount(0).
				Save(ctx)
			if err != nil {
				return errors.Wrap(err, "create stepRun")
			}

			if i == 0 && template.SelectionType == templateModel.SelectionTypeINTERNAL_DB {
				participants, err := r.filterParticipants(ctx, tx, template)
				if err != nil {
					return errors.Wrap(err, "filter participants")
				}

				for _, participant := range participants {
					_, err = participant.Update().AddSteps(stepRun).Save(ctx)
					if err != nil {
						return errors.Wrap(err, "add participant to stepRun")
					}
				}
			}
		}

		run, err = tx.Run.UpdateOne(run).
			SetStatus(runModel.StatusRUNNING).
			SetStartAt(time.Now()).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "create run")
		}

		return nil
	}); err != nil {
		log.Error().Err(err).Msg("start run: commit transaction")
	}

	r.startStep(run)
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
			// starting the first step

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
			// running subsequent step

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

func (r *Runtime) endStep(run *ent.Run, stepRun *ent.StepRun) {
	// TODO end anything MTurk is doing for this run.
	// On a HIT Step, we should make sure the HIT is deleted and people who
	// haven't  completed the HIT yet are not brought to the next round
}

func (r *Runtime) endRun(run *ent.Run) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := ent.WithTx(ctx, r.conn.Client, func(tx *ent.Tx) error {
		now := time.Now()

		currentRun, err := run.QueryCurrentStep().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "get last step")
		}

		_, err = tx.
			StepRun.
			UpdateOne(currentRun).
			SetStatus(steprunModel.StatusCREATED).
			SetEndedAt(now).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "update last step")
		}

		_, err = tx.
			Run.
			UpdateOne(run).
			SetStatus(runModel.StatusRUNNING).
			SetStartAt(now).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "update run")
		}

		return nil
	}); err != nil {
		log.Error().Err(err).Msg("end run: commit transaction")
	}
}
