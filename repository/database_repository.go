package repository

import (
	"context"
	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

//DatabaseRepository
//Implements the Repository interface's functions
type DatabaseRepository struct {
	DatabasePool *pgxpool.Pool
}

func NewDatabaseRepository(databasePool *pgxpool.Pool) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePool: databasePool,
	}
}

func (r *DatabaseRepository) CreateEvent(ctx context.Context, input *model.NewEvent) (*model.Event, error) {
	//TODO: implement me
	panic("implement me")
}

func (r *DatabaseRepository) UpdateEvent(ctx context.Context, id string, input *model.NewEvent) (*model.Event, error) {
	//TODO: implement me
	panic("implement me")
}

func (r *DatabaseRepository) DeleteEvent(ctx context.Context, id string) (*model.Event, error) {
	//TODO: implement me
	panic("implement me")
}

func (r *DatabaseRepository) GetEvent(ctx context.Context, id string) (*model.Event, error) {
	//TODO: implement me
	panic("implement me")
}
