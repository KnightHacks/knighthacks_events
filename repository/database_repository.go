package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"

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

	var event model.Event
	err := r.DatabasePool.QueryRow(ctx, "SELECT id, location, start_date, end_date, name, description FROM events WHERE id = $1", id).Scan(&event.ID, &event.Location,
		&event.StartDate, &event.EndDate, &event.Name, &event.Description)

	if err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &event, err

	//TODO: implement me
	//panic("implement me")
}

func (r *DatabaseRepository) UpdateEvent(ctx context.Context, id string, input *model.UpdatedEvent) (*model.Event, error) {
	//TODO implement me
	panic("implement me")
}
