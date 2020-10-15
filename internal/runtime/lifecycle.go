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

	nextStepRun, nextStep, err := r.getNextStep(ctx, run, currentStepRun)
	if err != nil {
		log.Error().Err(err).Msg("start step")
		return
	}

	now := time.Now()

	_, err = nextStepRun.
		Update().
		SetStatus(steprunModel.StatusRUNNING).
		SetStartedAt(now).
		Save(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("start step: update step run", run.ID)
		return
	}

	switch nextStep.Type {
	case stepModel.TypeMTURK_HIT, stepModel.TypeMTURK_MESSAGE:
		err = r.mturk.RunStep(run, nextStepRun, nextStep)
		if err != nil {
			log.Error().Err(err).Msg("start step: run mturk for step")
		}
	case stepModel.TypePARTICIPANT_FILTER:
		particpants, err := currentStepRun.QueryParticipants().All(ctx)
		if err != nil {
			log.Error().Err(err).Msg("start step: get participants for filter step failed")
			break
		}

		_, err = nextStepRun.Update().
			AddParticipants(particpants...).
			Save(ctx)
		if err != nil {
			log.Error().Err(err).Msg("start step: push filter step participants to next run")
		}

		log.Warn().Msg("filter step not implemented")
	default:
		log.Error().Err(err).Msgf("unknown step type: %s", nextStep.Type.String())
	}
}

func (r *Runtime) endStep(run *ent.Run, stepRun *ent.StepRun) {
	// TODO end anything MTurk is doing for this run.
	// On a HIT Step, we should make sure the HIT is stopped and people who
	// haven't completed the HIT yet are not brought to the next round
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := stepRun.
		Update().
		SetStatus(steprunModel.StatusDONE).
		SetEndedAt(time.Now()).
		Save(ctx)
	if err != nil {
		log.Error().Err(err).Msg("end run: update last step")
	}

	step, err := stepRun.QueryStep().Only(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("end step: get step (%s)", run.ID)
		return
	}

	nextStepRun, nextStep, err := r.getNextStep(ctx, run, stepRun)
	if err != nil {
		log.Error().Err(err).Msgf("end step: get next step (%s)", run.ID)
		return
	}

	switch nextStep.Type {
	case stepModel.TypeMTURK_HIT, stepModel.TypeMTURK_MESSAGE:
		err = r.mturk.EndStep(run, stepRun, step, nextStepRun, nextStep)
		if err != nil {
			log.Error().Err(err).Msg("start step: run mturk for step")
		}
	case stepModel.TypePARTICIPANT_FILTER:
		// noop?
	default:
		log.Error().Err(err).Msgf("end step: unknown step type: %s", nextStep.Type.String())
	}

}

func (r *Runtime) endRun(run *ent.Run) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := ent.WithTx(ctx, r.conn.Client, func(tx *ent.Tx) error {
		now := time.Now()

		currentStepRun, err := run.QueryCurrentStep().Only(ctx)
		if err != nil {
			return errors.Wrap(err, "end run: get last step")
		}

		if currentStepRun != nil {
			r.endStep(run, currentStepRun)
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

func (r *Runtime) getNextStep(ctx context.Context, run *ent.Run, currentStepRun *ent.StepRun) (*ent.StepRun, *ent.Step, error) {

	stepRuns, err := run.QuerySteps().WithStep().All(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, nil, errors.Wrap(err, "start step: get current step")
	}

	if len(stepRuns) == 0 {
		return nil, nil, errors.Wrapf(err, "trying to run a Run without Steps! (%s)", run.ID)
	}

	if currentStepRun != nil {
		r.endStep(run, currentStepRun)
	}

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
			return nil, nil, errors.Errorf("could not find first step for run (%s)", run.ID)
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
			return nil, nil, errors.Errorf("could not find next step for run (run: %s, prev step: %s)", run.ID, currentStepRun.ID)
		}
	}

	nextStep, err := nextStepRun.QueryStep().Only(ctx)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "get next step", run.ID)
	}

	return nextStepRun, nextStep, nil
}