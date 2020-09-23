package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
)

func (r *adminResolver) Email(ctx context.Context, obj *ent.Admin) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *procedureResolver) Creator(ctx context.Context, obj *ent.Procedure) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *procedureResolver) SelectionType(ctx context.Context, obj *ent.Procedure) (*model.SelectionType, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *procedureResolver) InternalCriteria(ctx context.Context, obj *ent.Procedure) (*model.InternalCriteria, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *procedureResolver) MturkCriteria(ctx context.Context, obj *ent.Procedure) (*model.MTurkCriteria, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *procedureResolver) Steps(ctx context.Context, obj *ent.Procedure) ([]*model.Step, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Creator(ctx context.Context, obj *ent.Project) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Procedures(ctx context.Context, obj *ent.Project) ([]*ent.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Runs(ctx context.Context, obj *ent.Project) ([]*ent.Run, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Data(ctx context.Context, obj *ent.Project, keys []string) ([]*model.Datum, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Creator(ctx context.Context, obj *ent.Run) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Procedure(ctx context.Context, obj *ent.Run) (*ent.Procedure, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Status(ctx context.Context, obj *ent.Run) (model.Status, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Steps(ctx context.Context, obj *ent.Run) ([]*ent.StepRun, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) CurrentStep(ctx context.Context, obj *ent.Run) (*ent.StepRun, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Data(ctx context.Context, obj *ent.Run, keys []string) ([]*model.Datum, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepRunResolver) Step(ctx context.Context, obj *ent.StepRun) (*model.Step, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepRunResolver) Status(ctx context.Context, obj *ent.StepRun) (model.Status, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepRunResolver) StartedAt(ctx context.Context, obj *ent.StepRun) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepRunResolver) Participants(ctx context.Context, obj *ent.StepRun, first *int, after *string) (*model.ParticipantsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

// Admin returns generated.AdminResolver implementation.
func (r *Resolver) Admin() generated.AdminResolver { return &adminResolver{r} }

// Procedure returns generated.ProcedureResolver implementation.
func (r *Resolver) Procedure() generated.ProcedureResolver { return &procedureResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Run returns generated.RunResolver implementation.
func (r *Resolver) Run() generated.RunResolver { return &runResolver{r} }

// StepRun returns generated.StepRunResolver implementation.
func (r *Resolver) StepRun() generated.StepRunResolver { return &stepRunResolver{r} }

type adminResolver struct{ *Resolver }
type procedureResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type runResolver struct{ *Resolver }
type stepRunResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *procedureResolver) ParticipantCount(ctx context.Context, obj *ent.Procedure) (*int, error) {
	panic(fmt.Errorf("not implemented"))
}
