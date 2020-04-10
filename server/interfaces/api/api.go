package api

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"net/http"
	"server/application/usercase"
	"server/interfaces/api/graph"
	"server/interfaces/api/graph/generated"
)

type Config struct {
	AuthUsecase   usecase.AuthUsecase
	TodoUsecase   usecase.TodoUsecase
	ToggleUsecase usecase.ToggleUsecase
	UserUsecase   usecase.UserUsecase
}

func NewHandler(config Config) http.Handler {
	r := chi.NewRouter()
	// GraphQL
	rootResolver := &graph.Resolver{
		AuthUsecase:   config.AuthUsecase,
		TodoUsecase:   config.TodoUsecase,
		ToggleUsecase: config.ToggleUsecase,
		UserUsecase:   config.UserUsecase,
	}
	schemaConfig := generated.Config{
		Resolvers:  rootResolver,
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}
	es := generated.NewExecutableSchema(schemaConfig)
	r.Handle("/", playground.Handler("GraphQL playground", "/api/query"))
	r.Handle("/query", handler.NewDefaultServer(es))

	return r
}
