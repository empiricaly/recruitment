package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/datum"
	"github.com/empiricaly/recruitment/internal/ent/predicate"
	"github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/ent/step"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	errs "github.com/pkg/errors"
)

func (r *datumResolver) Versions(ctx context.Context, obj *ent.Datum) ([]*ent.Datum, error) {
	p, err := obj.QueryParticipant().Only(ctx)
	if err != nil {
		return nil, errs.Wrap(err, "query participant")
	}

	return p.
		QueryData().
		Where(datum.And(datum.KeyEQ(obj.Key))).
		Order(ent.Asc(datum.FieldVersion), ent.Asc(datum.FieldIndex)).
		All(ctx)
}

func (r *filterStepArgsResolver) Type(ctx context.Context, obj *model.FilterStepArgs) (*model.ParticipantFilterType, error) {
	if obj.Type == nil {
		return nil, nil
	}
	t := model.ParticipantFilterType(obj.Type.String())
	return &t, nil
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

func (r *participantResolver) CreatedBy(ctx context.Context, obj *ent.Participant) (*ent.StepRun, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *participantResolver) Steps(ctx context.Context, obj *ent.Participant) ([]*ent.StepRun, error) {
	return obj.QuerySteps().Order(ent.Asc(step.FieldCreatedAt)).All(ctx)
}

func (r *participantResolver) ProviderIDs(ctx context.Context, obj *ent.Participant) ([]*ent.ProviderID, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *participantResolver) Data(ctx context.Context, obj *ent.Participant, keys []string, deleted *bool) ([]*ent.Datum, error) {
	predicates := []predicate.Datum{
		datum.Current(true),
	}

	if deleted != nil && *deleted {
		predicates = append(predicates, datum.DeletedAtNotNil())
	} else {
		predicates = append(predicates, datum.DeletedAtIsNil())
	}

	if keys != nil {
		predicates = append(predicates, datum.KeyIn(keys...))
	}

	return obj.QueryData().Where(datum.And(predicates...)).
		Order(ent.Asc(datum.FieldKey), ent.Asc(datum.FieldIndex)).
		All(ctx)
}

func (r *participationResolver) Step(ctx context.Context, obj *ent.Participation) (*ent.StepRun, error) {
	return obj.QueryStepRun().Only(ctx)
}

func (r *participationResolver) Participant(ctx context.Context, obj *ent.Participation) (*ent.Participant, error) {
	return obj.QueryParticipant().Only(ctx)
}

func (r *projectResolver) Creator(ctx context.Context, obj *ent.Project) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Templates(ctx context.Context, obj *ent.Project) ([]*ent.Template, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *projectResolver) Runs(ctx context.Context, obj *ent.Project, runID *string, statuses []model.Status, limit *int) ([]*ent.Run, error) {
	if runID != nil {
		return obj.QueryRuns().Where(run.IDEQ(*runID)).All(ctx)
	}

	q := obj.QueryRuns().Order(ent.Desc(run.FieldCreatedAt))
	if limit != nil && *limit > 0 {
		q = q.Limit(*limit)
	}

	if len(statuses) > 0 {
		s := make([]run.Status, len(statuses))
		for i, st := range statuses {
			s[i] = run.Status(st)
		}
		q = q.Where(run.StatusIn(s...))
	}

	return q.All(ctx)
}

func (r *projectResolver) Participants(ctx context.Context, obj *ent.Project, offset *int, limit *int) ([]*ent.Participant, error) {
	q := obj.QueryParticipants()
	if offset != nil && limit != nil && *limit > 0 {
		skip := *offset * *limit
		q = q.Limit(*limit).Offset(skip)
	}

	return q.All(ctx)
}

func (r *projectResolver) ParticipantsCount(ctx context.Context, obj *ent.Project) (int, error) {
	return obj.QueryParticipants().Count(ctx)
}

func (r *providerIDResolver) ProviderID(ctx context.Context, obj *ent.ProviderID) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *providerIDResolver) Provider(ctx context.Context, obj *ent.ProviderID) (*model.Provider, error) {
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
	return obj.QuerySteps().Order(ent.Asc(step.FieldIndex)).All(ctx)
}

func (r *runResolver) CurrentStep(ctx context.Context, obj *ent.Run) (*ent.StepRun, error) {
	return obj.QueryCurrentStep().Only(ctx)
}

func (r *stepResolver) Creator(ctx context.Context, obj *ent.Step) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepResolver) Type(ctx context.Context, obj *ent.Step) (model.StepType, error) {
	return model.StepType(obj.Type.String()), nil
}

func (r *stepRunResolver) Creator(ctx context.Context, obj *ent.StepRun) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepRunResolver) Step(ctx context.Context, obj *ent.StepRun) (*ent.Step, error) {
	return obj.QueryStep().Only(ctx)
}

func (r *stepRunResolver) Status(ctx context.Context, obj *ent.StepRun) (model.Status, error) {
	return model.Status(obj.Status.String()), nil
}

func (r *stepRunResolver) Participations(ctx context.Context, obj *ent.StepRun, first *int, after *string) ([]*ent.Participation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *stepRunResolver) Participants(ctx context.Context, obj *ent.StepRun, offset *int, limit *int) ([]*ent.Participant, error) {
	q := obj.QueryParticipants()
	if offset != nil && limit != nil && *limit > 0 {
		skip := *offset * *limit
		q = q.Limit(*limit).Offset(skip)
	}

	return q.All(ctx)
}

func (r *stepRunResolver) ParticipantsCount(ctx context.Context, obj *ent.StepRun) (int, error) {
	return obj.QueryParticipants().Count(ctx)
}

func (r *templateResolver) Creator(ctx context.Context, obj *ent.Template) (*ent.Admin, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *templateResolver) SelectionType(ctx context.Context, obj *ent.Template) (model.SelectionType, error) {
	return model.SelectionType(obj.SelectionType), nil
}

func (r *templateResolver) Steps(ctx context.Context, obj *ent.Template) ([]*ent.Step, error) {
	return obj.QuerySteps().Order(ent.Asc(step.FieldIndex)).All(ctx)
}

// Datum returns generated.DatumResolver implementation.
func (r *Resolver) Datum() generated.DatumResolver { return &datumResolver{r} }

// FilterStepArgs returns generated.FilterStepArgsResolver implementation.
func (r *Resolver) FilterStepArgs() generated.FilterStepArgsResolver {
	return &filterStepArgsResolver{r}
}

// MessageStepArgs returns generated.MessageStepArgsResolver implementation.
func (r *Resolver) MessageStepArgs() generated.MessageStepArgsResolver {
	return &messageStepArgsResolver{r}
}

// Participant returns generated.ParticipantResolver implementation.
func (r *Resolver) Participant() generated.ParticipantResolver { return &participantResolver{r} }

// Participation returns generated.ParticipationResolver implementation.
func (r *Resolver) Participation() generated.ParticipationResolver { return &participationResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// ProviderID returns generated.ProviderIDResolver implementation.
func (r *Resolver) ProviderID() generated.ProviderIDResolver { return &providerIDResolver{r} }

// Run returns generated.RunResolver implementation.
func (r *Resolver) Run() generated.RunResolver { return &runResolver{r} }

// Step returns generated.StepResolver implementation.
func (r *Resolver) Step() generated.StepResolver { return &stepResolver{r} }

// StepRun returns generated.StepRunResolver implementation.
func (r *Resolver) StepRun() generated.StepRunResolver { return &stepRunResolver{r} }

// Template returns generated.TemplateResolver implementation.
func (r *Resolver) Template() generated.TemplateResolver { return &templateResolver{r} }

type datumResolver struct{ *Resolver }
type filterStepArgsResolver struct{ *Resolver }
type messageStepArgsResolver struct{ *Resolver }
type participantResolver struct{ *Resolver }
type participationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type providerIDResolver struct{ *Resolver }
type runResolver struct{ *Resolver }
type stepResolver struct{ *Resolver }
type stepRunResolver struct{ *Resolver }
type templateResolver struct{ *Resolver }
