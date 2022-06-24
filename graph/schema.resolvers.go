package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/KnightHacks/knighthacks_events/graph/generated"
	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/KnightHacks/knighthacks_shared/auth"
	"github.com/KnightHacks/knighthacks_shared/models"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	return r.Repository.CreateEvent(ctx, &input)
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, input model.UpdatedEvent) (*model.Event, error) {
	return r.Repository.UpdateEvent(ctx, id, &input)
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	claims, ok := ctx.Value("AuthorizationUserClaims").(*auth.UserClaims)
	if !ok {
		return false, errors.New("unable to retrieve user claims, most likely forgot to set @hasRole directive")
	}
	if claims.Role != models.RoleAdmin {
		return false, errors.New("unauthorized to delete event")
	}

	return r.Repository.DeleteEvent(ctx, id)
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
