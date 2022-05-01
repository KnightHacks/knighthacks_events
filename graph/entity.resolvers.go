package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/LockedThread/knighthacks_events/graph/generated"
	"github.com/LockedThread/knighthacks_events/graph/model"
)

func (r *entityResolver) FindEventByID(ctx context.Context, id string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
