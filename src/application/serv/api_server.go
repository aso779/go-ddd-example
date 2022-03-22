package serv

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/aso779/go-ddd-example/application/config"
	"github.com/aso779/go-ddd-example/infrastructure/graph"
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
	conf *config.Config,
	log *zap.Logger,
	mux *chi.Mux,
	r *resolvers.Resolver,
) *APIServer {
	graphqlConfig := graph.Config{
		Resolvers: r,
	}
	//complexity.SetComplexityRules(&graphqlConfig.Complexity)
	queryHandler := handler.New(graph.NewExecutableSchema(graphqlConfig))
	queryHandler.AddTransport(transport.POST{})
	//queryHandler.Use(&extension.ComplexityLimit{Func: complexity.CalculateComplexity})
	queryHandler.Use(&extension.Introspection{})

	//queryHandler.Use(dataloaders.NewDataloaders(services, rel))
	//
	//queryHandler.SetErrorPresenter(handleErrors)
	//queryHandler.SetRecoverFunc(handleInternal)

	mux.
		//With(
		//	middlewares.UserMiddleware(services.User, services.UserJWT),
		//).
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
