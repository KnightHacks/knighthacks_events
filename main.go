package main

import (
	"context"
	"errors"
	"github.com/KnightHacks/knighthacks_events/graph/model"
	"github.com/KnightHacks/knighthacks_events/repository"
	"github.com/KnightHacks/knighthacks_shared/auth"
	"github.com/KnightHacks/knighthacks_shared/pagination"
	"github.com/KnightHacks/knighthacks_shared/utils"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KnightHacks/knighthacks_events/graph"
	"github.com/KnightHacks/knighthacks_events/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pool, err := pgxpool.Connect(context.Background(), utils.GetEnvOrDie("DATABASE_URI"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	databaseRepository := repository.NewDatabaseRepository(pool)

	hasRoleDirective := auth.HasRoleDirective{GetUserId: func(ctx context.Context, obj interface{}) (string, error) {
		switch _ := obj.(type) {
		case *model.Event:
			// TODO: Events don't have a sense of ownership, maybe we should have events linked to their user/sponsor hosting them?
			return "", errors.New("this shouldn't happen")
		default:
			// shouldn't happen, you must implement the new object with the ID field
			return "", errors.New("this shouldn't happen")
		}
	}}

	config := generated.Config{
		Resolvers: &graph.Resolver{
			Repository: databaseRepository,
		},
		Directives: generated.DirectiveRoot{
			HasRole:    hasRoleDirective.Direct,
			Pagination: pagination.Pagination,
		},
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
