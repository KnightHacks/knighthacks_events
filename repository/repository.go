package repository

import (
	"context"

	"github.com/KnightHacks/knighthacks_events/graph/model"
)

type Repository interface {
	CreateEvent(ctx context.Context, input *model.NewEvent) (*model.Event, error)
	UpdateEvent(ctx context.Context, id string, input *model.UpdatedEvent) (*model.Event, error)
	DeleteEvent(ctx context.Context, id string) (bool, error)
	GetEvent(ctx context.Context, id string) (*model.Event, error)
}
