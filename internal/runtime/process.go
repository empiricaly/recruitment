package runtime

import (
	"context"
	"math"
	"math/rand"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/participation"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	steprunModel "github.com/empiricaly/recruitment/internal/ent/steprun"
	templateModel "github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

//go:generate rice embed-go

func (r *runState) startRun(ctx context.Context, startTime time.Time) error {
	if err := ent.WithTx(ctx, r.conn.Client, func(tx *ent.Tx) error {
		for i, step := range r.steps {
			urlToken := randomString(30)
			log.Debug().Str("token", urlToken).Msg("Create URL Token")
			stepRun, err := tx.
				StepRun.
				Create().
				SetID(xid.New().String()).
				SetIndex(i).
				SetUrlToken(urlToken).
				SetRun(r.run).
				SetStep(step).
				SetStatus(steprunModel.StatusCREATED).
				SetParticipantsCount(0).
				Save(ctx)
			if err != nil {
				return errors.Wrap(err, "create stepRun")
			}

			if i == 0 && r.template.SelectionType == templateModel.SelectionTypeINTERNAL_DB {
				c := r.template.InternalCriteria
				participants, err := r.filterParticipants(ctx, tx, r.template.ParticipantCount, c.All, c.Uninitialized, c.Condition)
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

		_, err := tx.Run.UpdateOneID(r.run.ID).
			SetStatus(runModel.StatusRUNNING).
			SetStartedAt(startTime).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "create run")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "commit transaction")
	}

	if err := r.refresh(); err != nil {
		return errors.Wrap(err, "refresh after start run")
	}

	return r.setNextStep()
}

func (r *runState) advanceSteps(ctx context.Context, at time.Time) error {
	if r.currentStepRun != nil {
		_, err := r.currentStepRun.
			Update().
			SetStatus(steprunModel.StatusDONE).
			SetEndedAt(at).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "update previous step")
		}
	}

	if r.nextStep == nil {
		if r.endOfRunReached() {
			_, err := r.run.Update().
				ClearCurrentStep().
				Save(ctx)
			if err != nil {
				return errors.Wrap(err, "clear run's current step")
			}

			return r.refresh()
		} else {
			return errors.New("there is no next step defined")
		}
	}

	r.currentStep = r.nextStep
	r.currentStepRun = r.nextStepRun

	err := r.setNextStep()
	if err != nil {
		return errors.Wrap(err, "set next step")
	}

	_, err = r.currentStepRun.
		Update().
		SetStatus(steprunModel.StatusRUNNING).
		SetStartedAt(at).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "update new current step")
	}

	_, err = r.run.Update().
		SetCurrentStep(r.currentStepRun).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "update run's current step")
	}

	return r.refresh()
}

func (r *runState) startStep(ctx context.Context, startTime time.Time) error {
	switch r.currentStep.Type {
	case stepModel.TypeMTURK_HIT, stepModel.TypeMTURK_MESSAGE:
		err := r.mturkSession.StartStep(r.project, r.run, r.currentStep, r.currentStepRun, startTime)
		if err != nil {
			return errors.Wrap(err, "run mturk")
		}
	case stepModel.TypePARTICIPANT_FILTER:
		participants, err := r.currentStepRun.QueryParticipants().WithData().All(ctx)
		if err != nil {
			return errors.Wrap(err, "get participants for filter step failed")
		}

		if r.currentStep.FilterArgs == nil || r.currentStep.FilterArgs.Type == nil {
			log.Error().Msg("We are letting all participants through because the filterArgs object or type is missing")
		} else {
			switch *r.currentStep.FilterArgs.Type {
			case model.ParticipantFilterTypeJs:
				if r.currentStep.FilterArgs.Js == nil {
					log.Error().Msg("We are letting all participants through because the JS filter is empty")
					break
				}

				participants, err = jsfilter(ctx, r.conn, participants, *r.currentStep.FilterArgs.Js)
				if err != nil {
					return errors.Wrap(err, "jsfilter participants failed")
				}
			default:
				log.Error().Msgf("unsupported participant filter: %s", r.currentStep.FilterArgs.Type.String())
				log.Error().Msg("We are letting all participants through this unknown filter!")
			}
		}

		for _, p := range participants {
			prevParticipation, err := r.conn.Client.Participation.Query().Where(participation.MturkWorkerID(*p.MturkWorkerID)).First(ctx)
			if err != nil {
				log.Error().Err(err).Msg("Filter Participant: getting previous participation")
				return errors.Wrap(err, "filter step getting previous participation")
			}

			_, err = r.conn.Client.Participation.Create().
				SetID(xid.New().String()).
				SetParticipant(p).
				SetStepRun(r.currentStepRun).
				SetMturkWorkerID(*p.MturkWorkerID).
				SetMturkAssignmentID(prevParticipation.MturkAssignmentID).
				SetMturkHitID(prevParticipation.MturkHitID).
				Save(ctx)
			if err != nil {
				log.Error().Err(err).Msg("Filter Participant: Creating participation")
				return errors.Wrap(err, "filter step set participation")
			}
		}

	case stepModel.TypeWAIT:
		return nil
	default:
		return errors.Errorf("unknown step type: %s", r.currentStep.Type.String())
	}

	return r.refresh()
}

func (r *runState) endStep(ctx context.Context, endTime time.Time) error {
	switch r.currentStep.Type {
	case stepModel.TypeMTURK_HIT, stepModel.TypeMTURK_MESSAGE, stepModel.TypePARTICIPANT_FILTER:
		err := r.mturkSession.EndStep(r.project, r.run, r.currentStep, r.currentStepRun, r.nextStep, r.nextStepRun)
		if err != nil {
			return errors.Wrap(err, "run mturk for step")
		}
	case stepModel.TypeWAIT:
		return nil
	default:
		return errors.Errorf("unknown step type: %s", r.currentStep.Type.String())
	}

	return r.refresh()
}

func (r *runState) endRun(ctx context.Context, endTime time.Time) error {
	if err := ent.WithTx(ctx, r.conn.Client, func(tx *ent.Tx) error {
		_, err := tx.
			Run.
			UpdateOne(r.run).
			SetStatus(runModel.StatusDONE).
			SetEndedAt(endTime).
			Save(ctx)
		if err != nil {
			return errors.Wrap(err, "update run")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "commit transaction")
	}

	return r.refresh()
}

func (r *runState) filterParticipants(ctx context.Context, tx *ent.Tx, limit int, useAll bool, uninitialized bool, condition *model.Condition) ([]*ent.Participant, error) {
	participants, err := tx.Participant.Query().All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "find participants")
	}

	spew.Dump("BEFORE", participants)

	if !useAll {
		n := 0
		for _, participant := range participants {
			matches, err := participant.MatchCondition(ctx, condition)
			if err != nil {
				log.Debug().Err(err).Msg("match participant filter condition")
				continue
			}
			if matches {
				participants[n] = participant
				n++
			}
		}
		participants = participants[:n]
	}

	initializedParticipants := make([]*ent.Participant, 0)

	if uninitialized == true {
		initializedParticipants = participants
	} else {
		for _, participant := range participants {
			if participant.Uninitialized != nil && *participant.Uninitialized == true {
				continue
			}

			initializedParticipants = append(initializedParticipants, participant)
		}
	}

	spew.Dump("AFTER", initializedParticipants)

	// TODO should filter participants here
	l := math.Min(float64(limit), float64(len(initializedParticipants)))
	return initializedParticipants[:int(l)], nil
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// Generate a random string of A-Z chars with len = l
func randomString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		// 97, 122 to get lowercase
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}
