package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"strconv"
	"time"

	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	EventAlreadyExists = errors.New("event with id already exists")
	EventNotFound      = errors.New("event was not found")
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

/*
create table events
(
    id           serial,
    hackathon_id integer   not null,
    location     varchar   not null,
    start_date   timestamp not null,
    end_date     timestamp not null,
    name         varchar   not null,
    description  varchar   not null,
    constraint events_pk
        primary key (id),
    constraint events_hackathons_id_fk
        foreign key (hackathon_id) references hackathons
);
*/
func (r *DatabaseRepository) CreateEvent(ctx context.Context, input *model.NewEvent) (*model.Event, error) {
	//TODO: implement me
	var eventId string
	eventName := input.Name
	eventDescription := input.Description
	eventLocation := input.Location
	startDate := input.StartDate
	endDate := input.EndDate

	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		var discoveredEventID = new(int)
		//Okay, thinking about it it's possible that multiple events could have the same name at the same day.
		//Like HACKATHONA starts from 12:00 PM - 2:00 PM and then HACKATHONA has another event from 5:00 PM - 7:00 PM, but it's happening at some other location does something else
		//The real issue would be if you somehow try to make the same event on the same day twice.
		err := tx.QueryRow(ctx, "SELECT ID FROM events WHERE(NAME = $1) AND (START_DATE <= $3 AND END_DATE >= $2) AND (LOCATION = $4) LIMIT 1", eventName, startDate, endDate, eventLocation).Scan(&discoveredEventID)
		if err == nil && discoveredEventID != nil {
			return EventAlreadyExists
		}

		if err != nil && !errors.Is(err, pgx.ErrNoRows) {
			return err
		}

		var eventIdInt int
		//Not sure what happens with the hackathonid.
		err = tx.QueryRow(ctx, "INSERT INTO EVENTS (NAME,DESCRIPTION,START_DATE,END_DATE,LOCATION) VALUES ($1, $2, $3, $4, $5) RETURNING id",
			eventName,
			eventDescription,
			startDate,
			endDate,
			eventLocation).Scan(&eventIdInt)

		if err != nil {
			return err
		}
		eventId = strconv.Itoa(eventIdInt)
		return nil
	})

	if err != nil {
		return nil, err
	}

	//Hackathon id?
	return &model.Event{
		ID:          eventId,
		Location:    eventLocation,
		StartDate:   startDate,
		EndDate:     endDate,
		Name:        eventName,
		Description: eventDescription,
	}, nil
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
}

// UpdateEvent works where it checks to see if fields are nil or empty strings then it'll call the helper functions made
func (r *DatabaseRepository) UpdateEvent(ctx context.Context, id string, input *model.UpdatedEvent) (*model.Event, error) {
	var event model.Event
	if input.Name != nil && input.StartDate == nil && input.EndDate == nil && input.Description != nil && input.Location != nil {
		return nil, errors.New("empty event field")
	}
	err := r.DatabasePool.BeginTxFunc(ctx, pgx.TxOptions{}, func(tx pgx.Tx) error {
		if input.Name != nil {
			err := r.UpdateEventName(ctx, id, *input.Name, tx)
			if err != nil {
				return err
			}
			event.Name = *input.Name
		}
		if input.StartDate != nil {
			err := r.UpdateStartDate(ctx, id, *input.StartDate, tx)
			if err != nil {
				return err
			}
			event.StartDate = *input.StartDate
		}
		if input.EndDate != nil {
			err := r.UpdateEndDate(ctx, id, *input.EndDate, tx)
			if err != nil {
				return err
			}
			event.EndDate = *input.EndDate
		}
		if input.Description != nil {
			err := r.UpdateDescription(ctx, id, *input.Description, tx)
			if err != nil {
				return err
			}
			event.Description = *input.Description
		}
		if input.Location != nil {
			err := r.UpdateLocation(ctx, id, *input.Location, tx)
			if err != nil {
				return err
			}
			event.Location = *input.Location
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &event, nil
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
