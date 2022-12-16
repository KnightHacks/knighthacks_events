package integration_tests

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/KnightHacks/knighthacks_events/repository"
	"github.com/KnightHacks/knighthacks_shared/database"
	shared_db_utils "github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/utils"
	"github.com/jackc/pgx/v5/pgxpool"
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
		{
			name: "create Hackathon1",
			args: args{
				ctx: context.Background(),
				input: &model.NewEvent{
					Name:        "Hackathon1",
					StartDate:   time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
					EndDate:     time.Date(2000, time.February, 1, 1, 1, 1, 1, time.UTC),
					Description: "Hackathon1 Description",
					Location:    "UCF",
					HackathonID: "42",
				},
			},
			want: &model.Event{
				Name:        "Hackathon1",
				StartDate:   time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
				EndDate:     time.Date(2000, time.February, 1, 1, 1, 1, 1, time.UTC),
				Description: "Hackathon1 Description",
				Location:    "UCF",
			},
			wantErr: false,
		},
		// TODO: review
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
		{
			name: "delete Hackathon1",
			args: args{
				ctx: context.Background(),
				id:  "42",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "delete fake event",
			args: args{
				ctx: context.Background(),
				id:  "13579111315171921",
			},
			want:    false,
			wantErr: true,
		},
		// TODO: review
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
		{
			name: "get event 1",
			args: args{
				ctx: context.Background(),
				id:  "1",
			},
			want: &model.Event{
				ID:          "1",
				Name:        "event 1",
				StartDate:   time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
				EndDate:     time.Date(2000, time.February, 1, 1, 1, 1, 1, time.UTC),
				Description: "event 1 description",
				Location:    "event 1 location",
			},
		},
		// TODO: review
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
		{
			name: "get 5 events",
			args: args{
				ctx:   context.Background(),
				first: 5,   // what does this mean?
				after: "2", // what does this mean?
			},
			want: want{
				events: []*model.Event{
					{
						ID:          "2",
						Name:        "event 2",
						StartDate:   time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
						EndDate:     time.Date(2000, time.February, 1, 1, 1, 1, 1, time.UTC),
						Description: "event 2 description",
						Location:    "event 2 location",
					},
					{
						ID:          "3",
						Name:        "event 3",
						StartDate:   time.Date(2001, time.January, 1, 1, 1, 1, 1, time.UTC),
						EndDate:     time.Date(2001, time.February, 1, 1, 1, 1, 1, time.UTC),
						Description: "event 3 description",
						Location:    "event 3 location",
					},
					{
						ID:          "4",
						Name:        "event 4",
						StartDate:   time.Date(2002, time.January, 1, 1, 1, 1, 1, time.UTC),
						EndDate:     time.Date(2002, time.February, 1, 1, 1, 1, 1, time.UTC),
						Description: "event 4 description",
						Location:    "event 4 location",
					},
					{
						ID:          "5",
						Name:        "event 5",
						StartDate:   time.Date(2003, time.January, 1, 1, 1, 1, 1, time.UTC),
						EndDate:     time.Date(2003, time.February, 1, 1, 1, 1, 1, time.UTC),
						Description: "event 5 description",
						Location:    "event 5 location",
					},
					{
						ID:          "6",
						Name:        "event 6",
						StartDate:   time.Date(2004, time.January, 1, 1, 1, 1, 1, time.UTC),
						EndDate:     time.Date(2004, time.February, 1, 1, 1, 1, 1, time.UTC),
						Description: "event 6 description",
						Location:    "event 6 location",
					},
				},
				total: -1, // what does this mean?
			},
		},
		// TODO: review
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
		tx          database.Queryable
	}

	tests := []Test[args, any]{
		{
			name: "update event 1's description",
			args: args{
				ctx:         context.Background(),
				id:          "1",
				description: "1 updated",
				tx:          databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid",
			args: args{
				ctx:         context.Background(),
				id:          "-1",
				description: "update invalid description",
				tx:          databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		// TODO: review
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
		tx      database.Queryable
	}
	tests := []Test[args, any]{
		{
			name: "update event 1's end date",
			args: args{
				ctx:     context.Background(),
				id:      "1",
				endDate: time.Date(2000, time.November, 1, 1, 1, 1, 1, time.UTC),
				tx:      databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid event end date",
			args: args{
				ctx:     context.Background(),
				id:      "-1",
				endDate: time.Date(2000, time.November, 1, 1, 1, 1, 1, time.UTC),
				tx:      databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		// TODO: review
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
		{
			name: "update event 2",
			args: args{
				ctx: context.Background(),
				id:  "2",
				input: &model.UpdatedEvent{
					Name:        utils.Ptr("event 2022"),
					StartDate:   utils.Ptr(time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC)),
					EndDate:     utils.Ptr(time.Date(2000, time.February, 1, 1, 1, 1, 1, time.UTC)),
					Description: utils.Ptr("event 2 updated to be event 2022"),
					Location:    utils.Ptr("UCF"),
				},
			},
			want: &model.Event{
				ID:          "2",
				Name:        "event 2022",
				StartDate:   time.Date(2000, time.January, 1, 1, 1, 1, 1, time.UTC),
				EndDate:     time.Date(2000, time.February, 1, 1, 1, 1, 1, time.UTC),
				Description: "event 2 updated to be event 2022",
				Location:    "UCF",
			},
			wantErr: false,
		},
		{
			name: "update invalid",
			args: args{
				ctx: context.Background(),
				id:  "-2",
				input: &model.UpdatedEvent{
					Name:        utils.Ptr("abcdefg"),
					StartDate:   utils.Ptr(time.Date(2000, time.April, 1, 1, 1, 1, 1, time.UTC)),
					EndDate:     utils.Ptr(time.Date(2000, time.May, 1, 1, 1, 1, 1, time.UTC)),
					Description: utils.Ptr("update failed description"),
					Location:    utils.Ptr("UCF"),
				},
			},
			want:    nil,
			wantErr: true,
		},
		// TODO: review
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
		tx        database.Queryable
	}
	tests := []Test[args, any]{
		{
			name: "update event 5 to event 5000",
			args: args{
				ctx:       context.Background(),
				id:        "5",
				eventName: "5000",
				tx:        databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		{
			name: "update invalid event 6",
			args: args{
				ctx:       context.Background(),
				id:        "-6",
				eventName: "6000",
				tx:        databaseRepository.DatabasePool,
			},
			wantErr: true,
		},
		// TODO: review
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
		tx       database.Queryable
	}
	tests := []Test[args, any]{
		{
			name: "update location of event 1",
			args: args{
				ctx:      context.Background(),
				id:       "1",
				location: "Hawaii",
				tx:       databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		// TODO: review
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
		tx        database.Queryable
	}
	tests := []Test[args, any]{
		{
			name: "update event 1 start date",
			args: args{
				ctx:       context.Background(),
				id:        "1",
				startDate: time.Date(2001, time.January, 1, 1, 1, 1, 1, time.UTC),
				tx:        databaseRepository.DatabasePool,
			},
			wantErr: false,
		},
		// TODO: review
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
		queryable database.Queryable
	}
	tests := []Test[args, *model.Event]{
		{
			name: "query for event 1",
			args: args{
				ctx:       context.Background(),
				id:        "1",
				queryable: databaseRepository.DatabasePool,
			},
			want: &model.Event{
				ID:          "1",
				Name:        "event 1",
				StartDate:   time.Date(2001, time.January, 1, 1, 1, 1, 1, time.UTC),
				EndDate:     time.Date(2001, time.February, 1, 1, 1, 1, 1, time.UTC),
				Description: "event 1 description",
				Location:    "event 1 location",
			},
			wantErr: false,
		},
		{
			name: "query invalid ID: -1",
			args: args{
				ctx:       context.Background(),
				id:        "-1",
				queryable: databaseRepository.DatabasePool,
			},
			want:    nil,
			wantErr: true,
		},
		// TODO: review
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
		{
			name: "default",
			args: args{
				databasePool: databaseRepository.DatabasePool,
			},
			want: &repository.DatabaseRepository{
				DatabasePool: databaseRepository.DatabasePool,
			},
		},
		// TODO: review
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.NewDatabaseRepository(tt.args.databasePool); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDatabaseRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
