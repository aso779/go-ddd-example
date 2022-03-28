package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type BookOneCreateInput struct {
	GenreID     int64      `json:"genreId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       PriceInput `json:"price"`
}

func (r *BookOneCreateInput) ToEntity() *domain.Book {
	return &domain.Book{
		GenreID:     r.GenreID,
		Title:       r.Title,
		Description: r.Description,
		Price: domain.Price{
			Amount:   uint64(r.Price.Amount),
			Currency: r.Price.Currency,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
