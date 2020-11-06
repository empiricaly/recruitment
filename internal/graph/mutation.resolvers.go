package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	adminU "github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/admin"
	"github.com/empiricaly/recruitment/internal/ent/datum"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	stepModel "github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/ent/template"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	errs "github.com/pkg/errors"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

func (r *mutationResolver) UpdateDatum(ctx context.Context, input *model.UpdateDatumInput) (*ent.Datum, error) {
	if input.NodeType != nil && *input.NodeType != model.DatumNodeTypeParticipant {
		log.Warn().Msgf("UpdateDatum: unknown nodeType: %s", string(*input.NodeType))
		return nil, errs.New("unknown nodeType")
	}

	var newDatum *ent.Datum
	if err := ent.WithTx(ctx, r.Store.Client, func(tx *ent.Tx) error {
		p, err := tx.Participant.Get(ctx, input.NodeID)
		if err != nil {
			return errs.Wrap(err, "get participant")
		}

		var isAppend bool
		if input.IsAppend != nil && *input.IsAppend {
			isAppend = true
		}
		ent.SetDatum(ctx, tx, p, input.Key, input.Val, isAppend)
		if err != nil {
			return errs.Wrap(err, "insert new datum")
		}

		return nil
	}); err != nil {
		return nil, errs.Wrap(err, "update datum: commit transaction")
	}

	return newDatum, nil
}

func (r *mutationResolver) DeleteDatum(ctx context.Context, input *model.DeleteDatumInput) ([]*ent.Datum, error) {
	if input.NodeType != nil && *input.NodeType != model.DatumNodeTypeParticipant {
		log.Warn().Msgf("DeleteDatum: unknown nodeType: %s", string(*input.NodeType))
		return nil, errs.New("unknown nodeType")
	}

	var data []*ent.Datum
	if err := ent.WithTx(ctx, r.Store.Client, func(tx *ent.Tx) error {
		p, err := tx.Participant.Get(ctx, input.NodeID)
		if err != nil {
			return errs.Wrap(err, "get participant")
		}

		data, err = p.QueryData().
			Where(datum.And(datum.Current(true), datum.DeletedAtNotNil(), datum.KeyEQ(input.Key))).
			All(ctx)
		if err != nil && !ent.IsNotFound(err) {
			return errs.Wrap(err, "get data")
		}

		// If data with key does not exist, noop
		// NOTE(np): I don't know if returning an error helps?
		if len(data) == 0 {
			return nil
		}

		now := time.Now()

		ids := make([]string, len(data))
		for i, d := range data {
			ids[i] = d.ID
		}
		_, err = tx.Datum.
			Update().
			Where(datum.IDIn(ids...)).
			SetDeletedAt(now).
			Save(ctx)
		if err != nil {
			return errs.Wrap(err, "set deletedAt")
		}

		return nil
	}); err != nil {
		return nil, errs.Wrap(err, "delete datum: commit transaction")
	}

	return data, nil
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

func (r *mutationResolver) CreateTemplate(ctx context.Context, input *model.CreateTemplateInput) (*ent.Template, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTemplate(ctx context.Context, input *model.UpdateTemplateInput) (*ent.Template, error) {
	internalCriteria := model.InternalCriteriaFromInput(input.Template.InternalCriteria)
	mturkCriteria := model.MturkCriteriaFromInput(input.Template.MturkCriteria)

	tx, err := r.Store.Tx(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "starting a transaction")
	}

	template, err := tx.Template.UpdateOneID(*input.Template.ID).
		SetName(input.Template.Name).
		SetSelectionType(template.SelectionType(input.Template.SelectionType.String())).
		SetParticipantCount(input.Template.ParticipantCount).
		SetInternalCriteria(internalCriteria).
		SetMturkCriteria(mturkCriteria).
		SetAdult(input.Template.Adult).
		SetSandbox(input.Template.Sandbox).
		Save(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "create template")
	}

	steps, err := template.QuerySteps().All(ctx)
	newSteps := make([]string, len(input.Template.Steps))
	for _, step := range input.Template.Steps {
		hitArgs := model.HITStepArgsFromInput(step.HitArgs)
		msgArgs := model.MsgArgsFromInput(step.MsgArgs)
		filterArgs := model.FilterArgsFromInput(step.FilterArgs)

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
					SetType(stepModel.Type(step.Type.String())).
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
			SetType(stepModel.Type(step.Type.String())).
			SetIndex(step.Index).
			SetDuration(step.Duration).
			SetFilterArgs(filterArgs).
			SetHitArgs(hitArgs).
			SetMsgArgs(msgArgs).
			SetTemplate(template).
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

	return template, err
}

func (r *mutationResolver) DuplicateRun(ctx context.Context, input *model.DuplicateRunInput) (*ent.Run, error) {
	var result *ent.Run

	if err := ent.WithTx(ctx, r.Store.Client, func(tx *ent.Tx) error {
		projectID := input.ToProjectID

		run, err := tx.Run.Get(ctx, input.RunID)
		if err != nil {
			return errs.Wrap(err, "query duplicateRun for run")
		}

		if projectID == nil {
			project := run.QueryProject().OnlyIDX(ctx)
			projectID = &project
		}

		template, err := tx.Run.QueryTemplate(run).Only(ctx)
		if err != nil {
			log.Error().Err(err).Msg("start run: query template")
			return errs.Wrap(err, "query duplicateRun for template")
		}

		steps, err := template.QuerySteps().All(ctx)
		if err != nil {
			log.Error().Err(err).Msg("start run: query steps")
			return errs.Wrap(err, "query duplicateRun for steps")
		}

		creator := adminU.ForContext(ctx)

		newTemplate, err := tx.Template.Create().
			SetID(xid.New().String()).
			SetName(template.Name).
			SetSelectionType(template.SelectionType).
			SetParticipantCount(template.ParticipantCount).
			SetInternalCriteria(template.InternalCriteria).
			SetMturkCriteria(template.MturkCriteria).
			SetCreator(creator).
			SetSandbox(template.Sandbox).
			SetProjectID(*projectID).
			Save(ctx)
		if err != nil {
			return errs.Wrap(err, "create template")
		}

		for _, step := range steps {
			_, err := tx.Step.Create().
				SetID(xid.New().String()).
				SetType(stepModel.Type(step.Type.String())).
				SetIndex(step.Index).
				SetDuration(step.Duration).
				SetFilterArgs(step.FilterArgs).
				SetHitArgs(step.HitArgs).
				SetMsgArgs(step.MsgArgs).
				SetTemplate(newTemplate).
				Save(ctx)
			if err != nil {
				return errs.Wrap(err, "create step")
			}
		}

		newRun, err := tx.Run.Create().
			SetID(xid.New().String()).
			SetStatus(runModel.StatusCREATED).
			SetTemplate(newTemplate).
			SetName(run.Name + " - copy").
			SetProjectID(*projectID).
			Save(ctx)

		result = newRun
		return nil
	}); err != nil {
		return nil, errs.Wrap(err, "duplicate run: commit transaction")
	}

	return result, nil
}

func (r *mutationResolver) CreateRun(ctx context.Context, input *model.CreateRunInput) (*ent.Run, error) {
	tx, err := r.Store.Tx(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "starting a transaction")
	}

	creator := adminU.ForContext(ctx)

	internalCriteria := model.InternalCriteriaFromInput(input.Template.InternalCriteria)
	mturkCriteria := model.MturkCriteriaFromInput(input.Template.MturkCriteria)

	template, err := tx.Template.Create().
		SetID(xid.New().String()).
		SetName(input.Template.Name).
		SetSelectionType(template.SelectionType(input.Template.SelectionType.String())).
		SetParticipantCount(input.Template.ParticipantCount).
		SetInternalCriteria(internalCriteria).
		SetMturkCriteria(mturkCriteria).
		SetCreator(creator).
		SetProjectID(input.ProjectID).
		SetAdult(input.Template.Adult).
		SetSandbox(input.Template.Sandbox).
		Save(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "create template")
	}

	run, err := tx.Run.Create().
		SetID(xid.New().String()).
		SetStatus(runModel.StatusCREATED).
		SetTemplate(template).
		SetName(input.Template.Name).
		SetProjectID(input.ProjectID).
		Save(ctx)

	err = tx.Commit()
	if err != nil {
		return nil, errs.Wrap(err, "commit transaction")
	}

	return run, err
}

func (r *mutationResolver) UpdateRun(ctx context.Context, input *model.UpdateRunInput) (*ent.Run, error) {
	run, err := r.Store.Run.UpdateOneID(input.ID).
		SetName(input.Name).
		Save(ctx)

	return run, err
}

func (r *mutationResolver) ScheduleRun(ctx context.Context, input *model.ScheduleRunInput) (*ent.Run, error) {
	run, err := r.Store.Run.Get(ctx, input.ID)
	if err != nil {
		return nil, errs.Wrap(err, "sheduleRun: get run")
	}

	err = run.Validate()
	if err != nil {
		return nil, err
	}

	run, err = r.Store.Run.UpdateOneID(input.ID).
		SetStartAt(input.StartAt).
		Save(ctx)

	return run, err
}

func (r *mutationResolver) UnscheduleRun(ctx context.Context, input *model.UnscheduleRunInput) (*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StartRun(ctx context.Context, input *model.StartRunInput) (*ent.Run, error) {
	run, err := r.Store.Run.Get(ctx, input.ID)
	if err != nil {
		return nil, errs.Wrap(err, "startRun: get run")
	}

	err = run.Validate()
	if err != nil {
		return nil, err
	}

	run, err = r.Store.Run.UpdateOneID(input.ID).
		SetStatus(runModel.StatusRUNNING).
		Save(ctx)

	return run, err
}

func (r *mutationResolver) CancelRun(ctx context.Context, input *model.CancelRunInput) (*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
