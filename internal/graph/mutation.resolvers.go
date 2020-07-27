package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/rs/xid"
)

func (r *mutationResolver) RegisterParticipant(ctx context.Context, input *model.RegisterParticipantInput) (*model.Participant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MutateDatum(ctx context.Context, input *model.MutateDatumInput) (*model.Datum, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Auth(ctx context.Context, input *model.AuthInput) (*model.AuthResp, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreateProjectInput) (*model.Project, error) {
	for _, project := range r.projects {
		if project.ProjectID == input.ProjectID {
			return nil, errors.New("Project with ID already exists")
		}
	}

	project := &model.Project{
		ID:        xid.New().String(),
		Name:      input.Name,
		ProjectID: input.ProjectID,
	}

	return project, r.Mapping.Txn(func(t *storage.MappingTxn) error {
		return t.AddProject(project)
	})

	r.projects = append(r.projects, project)
	// panic(fmt.Errorf("not implemented"))
	return project, nil
}

func (r *mutationResolver) CreateProcedure(ctx context.Context, input *model.CreateProcedureInput) (*model.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateProcedure(ctx context.Context, input *model.UpdateProcedureInput) (*model.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStep(ctx context.Context, input *model.UpdateStepInput) (*model.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DuplicateProcedure(ctx context.Context, input *model.DuplicateProcedureInput) (*model.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRun(ctx context.Context, input *model.CreateRunInput) (*model.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ScheduleRun(ctx context.Context, input *model.ScheduleRunInput) (*model.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnscheduleRun(ctx context.Context, input *model.UnscheduleRunInput) (*model.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StartRun(ctx context.Context, input *model.StartRunInput) (*model.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CancelRun(ctx context.Context, input *model.CancelRunInput) (*model.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
