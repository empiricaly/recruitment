package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
)

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	panic(fmt.Errorf("not implemented"))
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
