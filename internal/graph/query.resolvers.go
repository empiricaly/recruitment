package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mturk"
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

func (r *queryResolver) MturkQualificationTypes(ctx context.Context) ([]*model.MTurkQulificationType, error) {
	var quals []*model.MTurkQulificationType
	params := &mturk.ListQualificationTypesInput{MustBeRequestable: aws.Bool(true), MustBeOwnedByCaller: aws.Bool(true)}

	err := r.MTurk.ListQualificationTypesPages(params, func(page *mturk.ListQualificationTypesOutput, lastPage bool) bool {
		for _, qual := range page.QualificationTypes {
			quals = append(quals, &model.MTurkQulificationType{ID: *qual.QualificationTypeId, Name: *qual.Name, Description: *qual.Description})
		}
		return true
	})

	if err != nil {
		fmt.Println("Got error calling listQualification Type:")
		fmt.Print(err.Error())
	}

	return quals, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
