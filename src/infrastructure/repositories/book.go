package repositories

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
)

type BookRepository struct {
	CrudRepository[domain.Book, bun.Tx]
}

func NewBook(
	connSet *connection.ConnSet,
	c metadata.EntityMetaContainer,
) *BookRepository {
	return &BookRepository{struct {
		ConnSet *connection.ConnSet
		Meta    metadata.Meta
	}{ConnSet: connSet, Meta: c.Get(domain.Book{}.Name())}}
}
