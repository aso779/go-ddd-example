package repositories

import (
	"context"
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
	}{ConnSet: connSet, Meta: c.Get(domain.Book{}.EntityName())}}
}

func (r BookRepository) AddAuthor(
	ctx context.Context,
	tx bun.IDB,
	bookId int,
	authorId int,
) error {
	values := map[string]any{
		"book_id":   bookId,
		"author_id": authorId,
	}
	_, err := tx.NewInsert().
		Model(&values).
		TableExpr("bks_books_authors").
		Exec(ctx)

	return err
}

func (r BookRepository) DeleteAuthors(
	ctx context.Context,
	tx bun.Tx,
	bookId int,
) error {
	_, err := tx.NewDelete().
		Where("book_id = ?", bookId).
		TableExpr("bks_books_authors").
		Exec(ctx)

	return err
}
