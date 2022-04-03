package repositories

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
)

type GenreRepository struct {
	CrudRepository[domain.Genre, bun.Tx]
}

func NewGenre(
	connSet *connection.ConnSet,
	c metadata.EntityMetaContainer,
) *GenreRepository {
	return &GenreRepository{struct {
		ConnSet *connection.ConnSet
		Meta    metadata.Meta
	}{ConnSet: connSet, Meta: c.Get(domain.Genre{}.Name())}}
}
