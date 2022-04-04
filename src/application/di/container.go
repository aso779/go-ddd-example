package di

import (
	"github.com/aso779/go-ddd-example/application/config"
	"github.com/aso779/go-ddd-example/application/serv"
	"github.com/aso779/go-ddd-example/domain/usecases"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd-example/infrastructure/repositories"
	"github.com/aso779/go-ddd-example/infrastructure/services"
	"github.com/aso779/go-ddd-example/presentation/resolvers"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/go-chi/chi/v5"
	"go.uber.org/dig"
)

func BuildContainer() *dig.Container {
	c := dig.New()
	_ = c.Provide(NewEntities, dig.As(new(metadata.EntityMetaContainer)))
	_ = c.Provide(config.NewConfig)
	_ = c.Provide(chi.NewRouter)
	_ = c.Provide(services.NewLogger)
	_ = c.Provide(resolvers.NewResolver)
	_ = c.Provide(connection.NewPGConnSet)
	_ = c.Provide(serv.NewAPIServer)
	_ = c.Provide(repositories.NewBook)
	_ = c.Provide(services.NewBook, dig.As(new(usecases.BookService)))
	_ = c.Provide(repositories.NewGenre)
	_ = c.Provide(services.NewGenre, dig.As(new(usecases.GenreService)))
	_ = c.Provide(repositories.NewAuthor)
	_ = c.Provide(services.NewAuthor, dig.As(new(usecases.AuthorService)))

	return c
}
