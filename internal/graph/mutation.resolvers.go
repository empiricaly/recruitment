package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"

	adminU "github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/run"
	stepy "github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (r *mutationResolver) RegisterParticipant(ctx context.Context, input *model.RegisterParticipantInput) (*model.Participant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MutateDatum(ctx context.Context, input *model.MutateDatumInput) (*model.Datum, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Auth(ctx context.Context, input *model.AuthInput) (*model.AuthResp, error) {
	for _, adminCred := range r.Admins {
		if adminCred.Username == input.Username && adminCred.Password == input.Password {
			user, err := r.Store.Admin.Query().Where(admin.UsernameEQ(adminCred.Username)).First(ctx)

			if err != nil {
				return nil, errs.Wrap(err, "get admin")
			}

			authResp := &model.AuthResp{}
			authResp.Token, err = adminU.CreateUserIDToken([]byte(r.SecretKey), user.ID)
			if err != nil {
				return nil, errs.Wrap(err, "create admin token")
			}
			return authResp, nil
		}
	}

	return nil, errs.New("invalid user")
}

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreateProjectInput) (*ent.Project, error) {
	return r.Store.Project.Create().
		SetID(xid.New().String()).
		SetName(input.Name).
		SetProjectID(input.ProjectID).
		Save(ctx)
}

func (r *mutationResolver) CreateProcedure(ctx context.Context, input *model.CreateProcedureInput) (*ent.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProcedure(ctx context.Context, input *model.UpdateProcedureInput) (*ent.Procedure, error) {
	internalCriteria, err := json.Marshal(input.Procedure.InternalCriteria)
	if err != nil {
		return nil, errs.Wrap(err, "encode internal criteria")
	}
	mturkCriteria, err := json.Marshal(input.Procedure.MturkCriteria)
	if err != nil {
		return nil, errs.Wrap(err, "encode mturk criteria")
	}

	tx, err := r.Store.Tx(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "starting a transaction")
	}

	procedure, err := tx.Procedure.UpdateOneID(*input.Procedure.ID).
		SetName(input.Procedure.Name).
		SetSelectionType(input.Procedure.SelectionType.String()).
		SetParticipantCount(input.Procedure.ParticipantCount).
		SetInternalCriteria(internalCriteria).
		SetMturkCriteria(mturkCriteria).
		Save(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "create procedure")
	}

	steps, err := procedure.QuerySteps().All(ctx)
	newSteps := make([]string, len(input.Procedure.Steps))
	for _, step := range input.Procedure.Steps {
		filterArgs, err := json.Marshal(step.FilterArgs)
		if err != nil {
			return nil, errs.Wrap(err, "encode filter args")
		}

		hitArgs, err := json.Marshal(step.HitArgs)
		if err != nil {
			return nil, errs.Wrap(err, "encode hit args")
		}

		msgArgs, err := json.Marshal(step.MsgArgs)
		if err != nil {
			return nil, errs.Wrap(err, "encode msg args")
		}

		if step.ID != nil {
			var found bool
			var existinStep *ent.Step
			for _, existinStep = range steps {
				if existinStep.ID == *step.ID {
					found = true
					break
				}
			}

			if found {
				step, err := existinStep.Update().
					SetType(stepy.Type(step.Type.String())).
					SetIndex(step.Index).
					SetDuration(step.Duration).
					SetFilterArgs(filterArgs).
					SetHitArgs(hitArgs).
					SetMsgArgs(msgArgs).
					Save(ctx)
				if err != nil {
					return nil, errs.Wrap(err, "update step")
				}
				newSteps = append(newSteps, step.ID)
				continue
			}
		}

		step, err := tx.Step.Create().
			SetID(xid.New().String()).
			SetType(stepy.Type(step.Type.String())).
			SetIndex(step.Index).
			SetDuration(step.Duration).
			SetFilterArgs(filterArgs).
			SetHitArgs(hitArgs).
			SetMsgArgs(msgArgs).
			SetProcedure(procedure).
			Save(ctx)
		if err != nil {
			return nil, errs.Wrap(err, "create step")
		}

		newSteps = append(newSteps, step.ID)
	}

LOOP:
	for _, previousStep := range steps {
		for _, newStepID := range newSteps {
			if newStepID == previousStep.ID {
				continue LOOP
			}
		}

		err := tx.Step.DeleteOne(previousStep).Exec(ctx)
		if err != nil {
			log.Error().Err(err).Msg("failed to delete old step")
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, errs.Wrap(err, "commit transaction")
	}

	return procedure, err
}

func (r *mutationResolver) DuplicateProcedure(ctx context.Context, input *model.DuplicateProcedureInput) (*ent.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRun(ctx context.Context, input *model.CreateRunInput) (*ent.Run, error) {
	creator := adminU.ForContext(ctx)

	internalCriteria, err := json.Marshal(input.Procedure.InternalCriteria)
	if err != nil {
		return nil, errs.Wrap(err, "encode internal criteria")
	}
	mturkCriteria, err := json.Marshal(input.Procedure.MturkCriteria)
	if err != nil {
		return nil, errs.Wrap(err, "encode mturk criteria")
	}

	procedure, err := r.Store.Procedure.Create().
		SetID(xid.New().String()).
		SetName(input.Procedure.Name).
		SetSelectionType(input.Procedure.SelectionType.String()).
		SetParticipantCount(input.Procedure.ParticipantCount).
		SetInternalCriteria(internalCriteria).
		SetMturkCriteria(mturkCriteria).
		SetCreator(creator).
		SetProjectID(input.ProjectID).
		Save(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "create procedure")
	}

	run, err := r.Store.Run.Create().
		SetID(xid.New().String()).
		SetStatus(run.StatusCREATED).
		SetProcedure(procedure).
		SetName(input.Procedure.Name).
		SetProjectID(input.ProjectID).
		Save(ctx)

	return run, err
}

func (r *mutationResolver) UpdateRun(ctx context.Context, input *model.UpdateRunInput) (*ent.Run, error) {
	run, err := r.Store.Run.UpdateOneID(input.ID).
		SetName(input.Name).
		Save(ctx)

	return run, err
}

func (r *mutationResolver) ScheduleRun(ctx context.Context, input *model.ScheduleRunInput) (*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnscheduleRun(ctx context.Context, input *model.UnscheduleRunInput) (*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StartRun(ctx context.Context, input *model.StartRunInput) (*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CancelRun(ctx context.Context, input *model.CancelRunInput) (*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
