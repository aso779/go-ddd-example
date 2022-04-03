package serv

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/aso779/go-ddd-example/infrastructure/graph"
	"github.com/aso779/go-ddd-example/infrastructure/graph/dataloaders"
	"github.com/aso779/go-ddd-example/infrastructure/services"
	"github.com/aso779/go-ddd-example/presentation/resolvers"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

var routes []string

type APIServer struct {
	MuxCopy chi.Mux
}

func NewAPIServer(
	log *zap.Logger,
	mux *chi.Mux,
	r *resolvers.Resolver,
	services services.ServiceContainer,
) *APIServer {
	graphqlConfig := graph.Config{
		Resolvers: r,
	}
	queryHandler := handler.New(graph.NewExecutableSchema(graphqlConfig))
	queryHandler.AddTransport(transport.POST{})
	queryHandler.Use(&extension.Introspection{})
	queryHandler.Use(dataloaders.NewDataloaders(services))

	mux.
		Method("POST", "/graphql", queryHandler)
	if err := chi.Walk(mux, walkFunc); err != nil {
		fmt.Printf("Logging err: %s\n", err.Error())
	}

	log.Info("mux", zap.Strings("routes", routes))

	return &APIServer{
		MuxCopy: *mux,
	}
}

func walkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	routes = append(routes, fmt.Sprintf("%s %s\n", method, route))
	return nil
}
