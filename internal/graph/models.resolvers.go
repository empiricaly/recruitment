package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/empiricaly/recruitment/internal/graph/generated"
	"github.com/empiricaly/recruitment/internal/model"
)

func (r *datumResolver) Creator(ctx context.Context, obj *model.Datum) (model.Creator, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *datumResolver) DeletedAt(ctx context.Context, obj *model.Datum) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *datumResolver) Next(ctx context.Context, obj *model.Datum) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *datumResolver) Root(ctx context.Context, obj *model.Datum) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *datumResolver) Versions(ctx context.Context, obj *model.Datum) ([]*model.Datum, error) {
	panic(fmt.Errorf("not implemented"))
}

// Datum returns generated.DatumResolver implementation.
func (r *Resolver) Datum() generated.DatumResolver { return &datumResolver{r} }

type datumResolver struct{ *Resolver }
