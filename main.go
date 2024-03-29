package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/KnightHacks/knighthacks_events/graph"
	"github.com/KnightHacks/knighthacks_events/graph/generated"
	"github.com/KnightHacks/knighthacks_events/repository"
	"github.com/KnightHacks/knighthacks_shared/auth"
	"github.com/KnightHacks/knighthacks_shared/database"
	"github.com/KnightHacks/knighthacks_shared/pagination"
	"github.com/KnightHacks/knighthacks_shared/utils"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const defaultPort = "8080"

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pool, err := database.ConnectWithRetries(utils.GetEnvOrDie("DATABASE_URI"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	newAuth, err := auth.NewAuthWithEnvironment()
	if err != nil {
		log.Fatalf("An error occured when trying to create an instance of Auth: %s\n", err)
	}

	ginRouter := gin.Default()
	ginRouter.Use(auth.AuthContextMiddleware(newAuth))
	ginRouter.Use(utils.GinContextMiddleware())

	ginRouter.POST("/query", graphqlHandler(newAuth, pool))
	ginRouter.GET("/", playgroundHandler())

	log.Fatal(ginRouter.Run(":" + port))
}

func graphqlHandler(a *auth.Auth, pool *pgxpool.Pool) gin.HandlerFunc {
	// TODO: Sponsor doesn't have a sense of ownership, maybe we should have sponsor linked users?

	hasRoleDirective := auth.HasRoleDirective{GetUserId: auth.DefaultGetUserId}

	config := generated.Config{
		Resolvers: &graph.Resolver{
			Repository: repository.NewDatabaseRepository(pool),
			Auth:       a,
		},
		Directives: generated.DirectiveRoot{
			HasRole:    hasRoleDirective.Direct,
			Pagination: pagination.Pagination,
		},
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))
	srv.SetRecoverFunc(func(ctx context.Context, iErr interface{}) error {
		err := fmt.Errorf("%v", iErr)
		log.Println(fmt.Sprintf("runtime error: %v\n", err))
		debug.PrintStack()
		return gqlerror.Errorf("Internal server error! Check logs for more details!")
	})
	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		log.Println("Error presented: ", err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})
	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
