package integration_tests

import (
	"context"
	"flag"
	"fmt"
	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/KnightHacks/knighthacks_events/repository"
	shared_db_utils "github.com/KnightHacks/knighthacks_shared/database"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"reflect"
	"testing"
	"time"
)

var integrationTest = flag.Bool("integration", false, "whether to run integration tests")
var databaseUri = flag.String("postgres-uri", "postgresql://postgres:test@localhost:5432/postgres", "postgres uri for running integration tests")

var databaseRepository *repository.DatabaseRepository

type Test[A any, T any] struct {
	name    string
	args    A
	want    T
	wantErr bool
}

func TestMain(t *testing.M) {
	flag.Parse()
	// check if integration testing is disabled
	if *integrationTest == false {
		return
	}

	// connect to database
	var err error
	pool, err := shared_db_utils.ConnectWithRetries(*databaseUri)
	fmt.Printf("connecting to database, pool=%v, err=%v\n", pool, err)
	if err != nil {
		log.Fatalf("unable to connect to database err=%v\n", err)
	}

	databaseRepository = repository.NewDatabaseRepository(pool)
	os.Exit(t.Run())
}

func TestDatabaseRepository_CreateEvent(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *model.NewEvent
	}
	tests := []Test[args, *model.Event]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.CreateEvent(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_DeleteEvent(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []Test[args, bool]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.DeleteEvent(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DeleteEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetEvent(t *testing.T) {
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []Test[args, *model.Event]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := databaseRepository.GetEvent(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_GetEvents(t *testing.T) {
	type args struct {
		ctx   context.Context
		first int
		after string
	}
	type want struct {
		events []*model.Event
		total  int
	}
	tests := []Test[args, want]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			events, total, err := databaseRepository.GetEvents(tt.args.ctx, tt.args.first, tt.args.after)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEvents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(events, tt.want.events) {
				t.Errorf("GetEvents() got = %v, want %v", events, tt.want.events)
			}
			if total != tt.want.total {
				t.Errorf("GetEvents() got1 = %v, want %v", total, tt.want.total)
			}
		})
	}
}

func TestDatabaseRepository_UpdateDescription(t *testing.T) {
	type args struct {
		ctx         context.Context
		id          string
		description string
		tx          pgx.Tx
	}
	tests := []Test[args, any]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := databaseRepository.UpdateDescription(tt.args.ctx, tt.args.id, tt.args.description, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateEndDate(t *testing.T) {
	type args struct {
		ctx     context.Context
		id      string
		endDate time.Time
		tx      pgx.Tx
	}
	tests := []Test[args, any]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := databaseRepository.UpdateEndDate(tt.args.ctx, tt.args.id, tt.args.endDate, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEndDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateEvent(t *testing.T) {
	type args struct {
		ctx   context.Context
		id    string
		input *model.UpdatedEvent
	}
	tests := []Test[args, *model.Event]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.UpdateEvent(tt.args.ctx, tt.args.id, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateEvent() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatabaseRepository_UpdateEventName(t *testing.T) {
	type args struct {
		ctx       context.Context
		id        string
		eventName string
		tx        pgx.Tx
	}
	tests := []Test[args, any]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := databaseRepository.UpdateEventName(tt.args.ctx, tt.args.id, tt.args.eventName, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateEventName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateLocation(t *testing.T) {
	type args struct {
		ctx      context.Context
		id       string
		location string
		tx       pgx.Tx
	}
	tests := []Test[args, any]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := databaseRepository.UpdateLocation(tt.args.ctx, tt.args.id, tt.args.location, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_UpdateStartDate(t *testing.T) {
	type args struct {
		ctx       context.Context
		id        string
		startDate time.Time
		tx        pgx.Tx
	}
	tests := []Test[args, any]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if err := databaseRepository.UpdateStartDate(tt.args.ctx, tt.args.id, tt.args.startDate, tt.args.tx); (err != nil) != tt.wantErr {
				t.Errorf("UpdateStartDate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDatabaseRepository_getEventWithQueryable(t *testing.T) {
	type args struct {
		ctx       context.Context
		id        string
		queryable shared_db_utils.Queryable
	}
	tests := []Test[args, *model.Event]{
		{},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := databaseRepository.GetEventWithQueryable(tt.args.ctx, tt.args.id, tt.args.queryable)
			if (err != nil) != tt.wantErr {
				t.Errorf("getEventWithQueryable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEventWithQueryable() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDatabaseRepository(t *testing.T) {
	type args struct {
		databasePool *pgxpool.Pool
	}
	tests := []struct {
		name string
		args args
		want *repository.DatabaseRepository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewDatabaseRepository(tt.args.databasePool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
