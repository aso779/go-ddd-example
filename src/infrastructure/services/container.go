package services

import (
	"github.com/aso779/go-ddd-example/domain/usecases"
	"go.uber.org/dig"
)

type ServiceContainer struct {
	dig.In

	Book   usecases.BookService
	Genre  usecases.GenreService
	Author usecases.AuthorService
}
