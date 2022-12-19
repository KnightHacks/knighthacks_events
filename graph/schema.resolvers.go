package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/KnightHacks/knighthacks_events/graph/generated"
	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/KnightHacks/knighthacks_shared/pagination"
)

func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	return r.Repository.CreateEvent(ctx, &input)
}

func (r *mutationResolver) UpdateEvent(ctx context.Context, id string, input model.UpdatedEvent) (*model.Event, error) {
	return r.Repository.UpdateEvent(ctx, id, &input)
}

func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	return r.Repository.DeleteEvent(ctx, id)
}

func (r *queryResolver) Events(ctx context.Context, first int, after *string) (*model.EventsConnection, error) {
	a, err := pagination.DecodeCursor(after)
	if err != nil {
		return nil, err
	}
	events, total, err := r.Repository.GetEvents(ctx, first, a)
	if err != nil {
		return nil, err
	}

	return &model.EventsConnection{
		TotalCount: total,
		PageInfo:   pagination.GetPageInfo(events[0].ID, events[len(events)-1].ID),
		Events:     events,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
