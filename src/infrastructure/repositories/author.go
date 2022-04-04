package repositories

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/domain/projections"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
)

type AuthorRepository struct {
	CrudRepository[domain.Author, bun.Tx]
}

func NewAuthor(
	connSet *connection.ConnSet,
	c metadata.EntityMetaContainer,
) *AuthorRepository {
	return &AuthorRepository{struct {
		ConnSet *connection.ConnSet
		Meta    metadata.Meta
	}{ConnSet: connSet, Meta: c.Get(domain.Author{}.EntityName())}}
}

func (r *AuthorRepository) FindAllViaBookIds(
	ctx context.Context,
	fields []string,
	booksIDs []int,
) (*[]projections.AuthorBookID, error) {
	var ents []projections.AuthorBookID

	err := r.ConnSet.ReadPool().NewSelect().
		Model(&ents).
		Column(fields...).
		ColumnExpr("a.book_id").
		Join("JOIN bks_books_authors AS a ON a.author_id = bks_authors.id").
		Where("book_id IN (?)", bun.In(booksIDs)).
		Scan(ctx)

	return &ents, err
}
