package repository

import (
	"context"
	"errors"
	"github.com/KnightHacks/knighthacks_shared/database"
	"time"

	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	EventNotFound = errors.New("event was not found")
)

// DatabaseRepository
// Implements the Repository interface's functions
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
		return false, EventNotFound
	}

	// if the above conditions dont execute everything is good
	return true, nil
}

func (r *DatabaseRepository) GetEvent(ctx context.Context, id string) (*model.Event, error) {
	return r.getEventWithQueryable(ctx, id, r.DatabasePool)
}

func (r *DatabaseRepository) getEventWithQueryable(ctx context.Context, id string, queryable database.Queryable) (*model.Event, error) {
	var event model.Event
	err := queryable.QueryRow(ctx, "SELECT id, location, start_date, end_date, name, description FROM events WHERE id = $1", id).Scan(&event.ID, &event.Location,
		&event.StartDate, &event.EndDate, &event.Name, &event.Description)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, EventNotFound
		}
		return nil, err
	}

	return &event, err
}

// UpdateEvent works where it checks to see if fields are nil or empty strings then it'll call the helper functions made
func (r *DatabaseRepository) UpdateEvent(ctx context.Context, id string, input *model.UpdatedEvent) (*model.Event, error) {
	if input.Name == nil && input.StartDate == nil && input.EndDate == nil && input.Description == nil && input.Location == nil {
		return nil, errors.New("empty event field")
	}
	var event *model.Event
	var err error
	err = r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if input.Name != nil {
			err = r.UpdateEventName(ctx, id, *input.Name, tx)
			if err != nil {
				return err
			}
		}
		if input.StartDate != nil {
			err = r.UpdateStartDate(ctx, id, *input.StartDate, tx)
			if err != nil {
				return err
			}
		}
		if input.EndDate != nil {
			err = r.UpdateEndDate(ctx, id, *input.EndDate, tx)
			if err != nil {
				return err
			}
		}
		if input.Description != nil {
			err := r.UpdateDescription(ctx, id, *input.Description, tx)
			if err != nil {
				return err
			}
		}
		if input.Location != nil {
			err := r.UpdateLocation(ctx, id, *input.Location, tx)
			if err != nil {
				return err
			}
		}
		event, err = r.getEventWithQueryable(ctx, id, tx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *DatabaseRepository) UpdateEventName(ctx context.Context, id string, eventName string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET name = $1 WHERE id = $2", eventName, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return EventNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateStartDate(ctx context.Context, id string, startDate time.Time, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET start_date = $1 WHERE id = $2", startDate, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return EventNotFound
	}
	return nil
}

func (r *DatabaseRepository) UpdateEndDate(ctx context.Context, id string, endDate time.Time, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET end_date = $1 WHERE id = $2", endDate, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return EventNotFound
	}
	return nil
}
func (r *DatabaseRepository) UpdateDescription(ctx context.Context, id string, description string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET description = $1 WHERE id = $2", description, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return EventNotFound
	}
	return nil
}
func (r *DatabaseRepository) UpdateLocation(ctx context.Context, id string, location string, tx pgx.Tx) error {
	commandTag, err := tx.Exec(ctx, "UPDATE events SET location = $1 WHERE id = $2", location, id)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() != 1 {
		return EventNotFound
	}
	return nil
}

func (r *DatabaseRepository) GetEvents(ctx context.Context, first int, after string) ([]*model.Event, int, error) {
	events := make([]*model.Event, 0, first)
	var total int
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		rows, err := r.DatabasePool.Query(ctx, "SELECT id, location, start_date, end_date, name, description FROM events WHERE id > $1 ORDER BY `id` DESC LIMIT $2", after, first)
		if err != nil {
			return err
		}

		for rows.Next() {
			var event model.Event

			if err = rows.Scan(&event.ID, &event.Location, &event.StartDate, &event.EndDate, &event.Name, &event.Description); err != nil {
				return err
			}
			events = append(events, &event)
		}

		return r.DatabasePool.QueryRow(ctx, "SELECT COUNT(*) FROM events").Scan(&total)
	})

	if err != nil {
		return nil, 0, err
	}

	return events, total, nil
}
