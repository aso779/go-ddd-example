package dataloaders

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/aso779/go-ddd-example/infrastructure/services"
	"github.com/aso779/go-ddd-example/presentation/adapters"
	"time"
)

const loadersKey = "dataloaders"

type Loaders struct {
	GenreByGenreId Loader[adapters.GenreOutput]
}

type Dataloaders struct {
	services services.ServiceContainer
}

func NewDataloaders(
	s services.ServiceContainer,
) *Dataloaders {
	return &Dataloaders{
		services: s,
	}
}

func (r Dataloaders) ExtensionName() string {
	return "Dataloaders"
}

var _ interface {
	graphql.OperationInterceptor
	graphql.HandlerExtension
} = Dataloaders{}

func (r Dataloaders) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

func (r Dataloaders) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	nextCtx := context.WithValue(ctx, loadersKey, &Loaders{
		GenreByGenreId: Loader[adapters.GenreOutput]{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []int, fields []string) ([]*adapters.GenreOutput, []error) {
				//db query
				res, err := r.services.Genre.FindAllByIds(ctx, fields, ids)
				if err != nil {
					return nil, []error{err}
				}
				//map
				groupByGenreId := make(map[int]*adapters.GenreOutput, len(ids))
				for _, v := range *res {
					groupByGenreId[v.ID] = adapters.NewGenre().ToOutput(&v)
				}
				// order
				result := make([]*adapters.GenreOutput, len(ids))
				for i, id := range ids {
					result[i] = groupByGenreId[id]
				}

				return result, nil
			},
		},
	})

	return next(nextCtx)
}
