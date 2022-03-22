package domain

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
	"time"
)

type Book struct {
	bun.BaseModel `bun:"table:bks_books,alias:bks_books"`
	ID            int64  `bun:"id,pk,autoincrement" json:"id"`
	Title         string `bun:"title" json:"title"`
	//GenreID       int       `bun:"genre_id" json:"genreId"`
	CreatedAt time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at" json:"updatedAt"`
}

func (r Book) Name() string {
	return "Book"
}

func (r Book) PrimaryKey() metadata.PrimaryKey {
	return metadata.PrimaryKey{"id": r.ID}
}
