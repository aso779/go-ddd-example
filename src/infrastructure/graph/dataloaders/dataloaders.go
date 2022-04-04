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
	BookGenreByGenreId  LoaderOne[adapters.GenreOutput]
	BookAuthorsByBookId LoaderMany[adapters.AuthorOutput]
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
		BookGenreByGenreId: LoaderOne[adapters.GenreOutput]{
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
		BookAuthorsByBookId: LoaderMany[adapters.AuthorOutput]{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []int, fields []string) ([][]adapters.AuthorOutput, []error) {
				// db query
				res, err := r.services.Author.FindAllViaBookIds(ctx, fields, ids)
				if err != nil {
					return nil, []error{err}
				}
				// map
				groupByBookId := make(map[int][]adapters.AuthorOutput, len(ids))
				for _, v := range *res {
					groupByBookId[v.BookID] = append(
						groupByBookId[v.BookID],
						*adapters.NewAuthor().FromProjectionBookId(&v),
					)
				}
				// order
				result := make([][]adapters.AuthorOutput, len(ids))
				for i, bookId := range ids {
					result[i] = groupByBookId[bookId]
				}

				return result, nil
			},
		},
	})

	return next(nextCtx)
}
