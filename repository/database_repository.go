package repository

import (
	"context"
	"errors"

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

func (r *DatabaseRepository) DeleteEvent(ctx context.Context, id string) (bool, error) {

	// removes event
	commandTag, err := r.DatabasePool.Exec(ctx, "DELETE FROM events WHERE id = $1", id)

	// checks if there is an error
	if err != nil {
		return false, err
	}
	// checking to see if there is 1 row affected for deleted events if not there is an issue
	if commandTag.RowsAffected() != 1 {
		return false, errors.New("event was not found")
	}

	// if the above conditions dont execute everything is good
	return true, nil
}

func (r *DatabaseRepository) GetEvent(ctx context.Context, id string) (*model.Event, error) {
	//TODO: implement me
	panic("implement me")
}

func (r *DatabaseRepository) UpdateEvent(ctx context.Context, id string, input *model.UpdatedEvent) (*model.Event, error) {
	//TODO implement me
	panic("implement me")
}
