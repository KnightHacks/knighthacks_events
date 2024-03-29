package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/KnightHacks/knighthacks_events/graph/generated"
	"github.com/KnightHacks/knighthacks_events/graph/model"
)

// FindEventByID is the resolver for the findEventByID field.
func (r *entityResolver) FindEventByID(ctx context.Context, id string) (*model.Event, error) {
	return r.Repository.GetEvent(ctx, id)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
