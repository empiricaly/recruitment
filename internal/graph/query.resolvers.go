package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/empiricaly/recruitment/internal/admin"
	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/project"
	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
	"github.com/empiricaly/recruitment/internal/mturk"
	errs "github.com/pkg/errors"
)

func (r *participantsConnectionResolver) Participants(ctx context.Context, obj *model.ParticipantsConnection) ([]*ent.Participant, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *participationsConnectionResolver) Participations(ctx context.Context, obj *model.ParticipationsConnection) ([]*ent.Participation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Projects(ctx context.Context) ([]*ent.Project, error) {
	return r.Store.Project.Query().Order(ent.Desc(project.FieldCreatedAt)).All(ctx)
}

func (r *queryResolver) Project(ctx context.Context, id *string, projectID *string) (*ent.Project, error) {
	if id == nil && projectID == nil {
		return nil, errors.New("id or projectID required")
	}

	if id != nil {
		return r.Store.Project.Query().Where(project.IDEQ(*id)).First(ctx)
	}

	return r.Store.Project.Query().Where(project.ProjectIDEQ(*projectID)).First(ctx)
}

func (r *queryResolver) Participants(ctx context.Context, first *int, after *string) (*model.ParticipantsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (model.User, error) {
	return admin.ForContext(ctx), nil
}

func (r *queryResolver) Page(ctx context.Context, token string, participantID string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MturkQualificationTypes(ctx context.Context, sandbox *bool) ([]*model.MTurkQulificationType, error) {
	var mturkSession *mturk.Session
	if sandbox != nil && *sandbox {
		mturkSession = r.MTurkSanbox
	} else {
		mturkSession = r.MTurk
	}

	res, err := mturkSession.GetQuals()

	if err != nil {
		return nil, errs.Wrap(err, "get  qualificationTypes")
	}
	return res, nil
}

func (r *queryResolver) MturkLocales(ctx context.Context, sandbox *bool) ([]*model.MTurkLocale, error) {
	var mturkSession *mturk.Session
	if sandbox != nil && *sandbox {
		mturkSession = r.MTurkSanbox
	} else {
		mturkSession = r.MTurk
	}

	res, err := mturkSession.GetLocales()

	if err != nil {
		return nil, errs.Wrap(err, "get  locales")
	}
	return res, nil
}

// ParticipantsConnection returns generated.ParticipantsConnectionResolver implementation.
func (r *Resolver) ParticipantsConnection() generated.ParticipantsConnectionResolver {
	return &participantsConnectionResolver{r}
}

// ParticipationsConnection returns generated.ParticipationsConnectionResolver implementation.
func (r *Resolver) ParticipationsConnection() generated.ParticipationsConnectionResolver {
	return &participationsConnectionResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type participantsConnectionResolver struct{ *Resolver }
type participationsConnectionResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
