package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
)

func (r *filterStepArgsResolver) Type(ctx context.Context, obj *model.FilterStepArgs) (model.ParticipantFilterType, error) {
	return model.ParticipantFilterType(obj.Type.String()), nil
}

func (r *messageStepArgsResolver) MessageType(ctx context.Context, obj *model.MessageStepArgs) (model.ContentType, error) {
	return model.ContentType(obj.MessageType.String()), nil
}

func (r *messageStepArgsResolver) LobbyType(ctx context.Context, obj *model.MessageStepArgs) (*model.ContentType, error) {
	if obj.LobbyType == nil {
		return nil, nil
	}
	t := model.ContentType(obj.LobbyType.String())
	return &t, nil
}

func (r *projectResolver) Creator(ctx context.Context, obj *ent.Project) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Templates(ctx context.Context, obj *ent.Project) ([]*ent.Template, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Runs(ctx context.Context, obj *ent.Project, runID *string, limit *int) ([]*ent.Run, error) {
	if runID != nil {
		return obj.QueryRuns().Where(run.IDEQ(*runID)).All(ctx)
	}

	q := obj.QueryRuns().Order(ent.Desc(run.FieldCreatedAt))
	if limit != nil && *limit > 0 {
		q = q.Limit(*limit)
	}

	return q.All(ctx)
}

func (r *projectResolver) Data(ctx context.Context, obj *ent.Project, keys []string) ([]*model.Datum, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Creator(ctx context.Context, obj *ent.Run) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *runResolver) Template(ctx context.Context, obj *ent.Run) (*ent.Template, error) {
	return obj.QueryTemplate().First(ctx)
}

func (r *runResolver) Status(ctx context.Context, obj *ent.Run) (model.Status, error) {
	return model.Status(obj.Status.String()), nil
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

func (r *stepResolver) Creator(ctx context.Context, obj *ent.Step) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepResolver) Type(ctx context.Context, obj *ent.Step) (model.StepType, error) {
	return model.StepType(obj.Type.String()), nil
}

func (r *stepResolver) MsgArgs(ctx context.Context, obj *ent.Step) (*model.MessageStepArgs, error) {
	args := &model.MessageStepArgs{}
	err := json.Unmarshal(obj.MsgArgs, args)
	return args, err
}

func (r *stepResolver) HitArgs(ctx context.Context, obj *ent.Step) (*model.HITStepArgs, error) {
	args := &model.HITStepArgs{}
	err := json.Unmarshal(obj.HitArgs, args)
	return args, err
}

func (r *stepResolver) FilterArgs(ctx context.Context, obj *ent.Step) (*model.FilterStepArgs, error) {
	args := &model.FilterStepArgs{}
	err := json.Unmarshal(obj.FilterArgs, args)
	return args, err
}

func (r *stepRunResolver) Step(ctx context.Context, obj *ent.StepRun) (*ent.Step, error) {
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

func (r *templateResolver) Creator(ctx context.Context, obj *ent.Template) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *templateResolver) SelectionType(ctx context.Context, obj *ent.Template) (model.SelectionType, error) {
	return model.SelectionType(obj.SelectionType), nil
}

func (r *templateResolver) InternalCriteria(ctx context.Context, obj *ent.Template) (*model.InternalCriteria, error) {
	crit := &model.InternalCriteria{}
	err := json.Unmarshal(obj.InternalCriteria, crit)
	return crit, err
}

func (r *templateResolver) MturkCriteria(ctx context.Context, obj *ent.Template) (*model.MTurkCriteria, error) {
	crit := &model.MTurkCriteria{}
	err := json.Unmarshal(obj.MturkCriteria, crit)
	return crit, err
}

func (r *templateResolver) Steps(ctx context.Context, obj *ent.Template) ([]*ent.Step, error) {
	return obj.QuerySteps().Order(ent.Asc(step.FieldIndex)).All(ctx)
}

// FilterStepArgs returns generated.FilterStepArgsResolver implementation.
func (r *Resolver) FilterStepArgs() generated.FilterStepArgsResolver {
	return &filterStepArgsResolver{r}
}

// MessageStepArgs returns generated.MessageStepArgsResolver implementation.
func (r *Resolver) MessageStepArgs() generated.MessageStepArgsResolver {
	return &messageStepArgsResolver{r}
}

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Run returns generated.RunResolver implementation.
func (r *Resolver) Run() generated.RunResolver { return &runResolver{r} }

// Step returns generated.StepResolver implementation.
func (r *Resolver) Step() generated.StepResolver { return &stepResolver{r} }

// StepRun returns generated.StepRunResolver implementation.
func (r *Resolver) StepRun() generated.StepRunResolver { return &stepRunResolver{r} }

// Template returns generated.TemplateResolver implementation.
func (r *Resolver) Template() generated.TemplateResolver { return &templateResolver{r} }

type filterStepArgsResolver struct{ *Resolver }
type messageStepArgsResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type runResolver struct{ *Resolver }
type stepResolver struct{ *Resolver }
type stepRunResolver struct{ *Resolver }
type templateResolver struct{ *Resolver }
