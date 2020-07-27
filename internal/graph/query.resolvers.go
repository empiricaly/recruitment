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
	errs "github.com/pkg/errors"
)

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var err error
	var projects []*model.Project
	err = r.Mapping.Txn(func(t *storage.MappingTxn) error {
		projects, err = t.Projects()
		return err
	})
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *queryResolver) Project(ctx context.Context, id *string, projectID *string) (*model.Project, error) {
	if id == nil && projectID == nil {
		return nil, errors.New("id or projectID required")
	}
	var project *model.Project
	var err error
	err = r.Mapping.Txn(func(t *storage.MappingTxn) error {
		if projectID != nil {
			project, err = t.ProjectByProjectID(*projectID)
		} else if id != nil {
			project, err = t.Project(*id)
		}
		return err
	})

	return project, err
}

func (r *queryResolver) Participants(ctx context.Context, first *int, after *string) (*model.ParticipantsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Page(ctx context.Context, token string, participantID string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) MturkQualificationTypes(ctx context.Context) ([]*model.MTurkQulificationType, error) {
	res, err := r.MTurk.GetQuals()

	if err != nil {
		return nil, errs.Wrap(err, "get  qualificationTypes")
	}
	return res, nil
}

func (r *queryResolver) MturkLocales(ctx context.Context) ([]*model.MTurkLocale, error) {
	res, err := r.MTurk.GetLocales()

	if err != nil {
		return nil, errs.Wrap(err, "get  locales")
	}
	return res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
