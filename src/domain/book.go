package domain

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
	"time"
)

type Book struct {
	bun.BaseModel `bun:"table:bks_books,alias:bks_books"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	GenreID       int       `bun:"genre_id" json:"genreId"`
	Title         string    `bun:"title" json:"title"`
	Description   string    `bun:"description" json:"description"`
	Price         Price     `bun:"embed:price_"`
	CreatedAt     time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `bun:"updated_at" json:"updatedAt"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

type Price struct {
	Amount   uint   `bun:"amount" json:"amount"`
	Currency string `bun:"currency" json:"currency"`
}

func (r Book) Name() string {
	return "Book"
}

func (r Book) PrimaryKey() metadata.PrimaryKey {
	return metadata.PrimaryKey{"id": r.ID}
}

func (r *Book) ToExistsEntity(exists *Book) {
	if r.GenreID != 0 {
		exists.GenreID = r.GenreID
	}
	if r.Title != "" {
		exists.Title = r.Title
	}
	if r.Description != "" {
		exists.Description = r.Description
	}
	exists.Price.Amount = r.Price.Amount
	exists.Price.Currency = r.Price.Currency
}
