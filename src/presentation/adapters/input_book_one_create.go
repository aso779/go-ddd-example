package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/domain/values"
	"time"
)

type BookOneCreateInput struct {
	GenreID     int        `json:"genreId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       PriceInput `json:"price"`
	Authors     []int      `json:"authors"`
}

func (r *BookOneCreateInput) ToEntity() *domain.Book {
	return &domain.Book{
		GenreID:     r.GenreID,
		Title:       r.Title,
		Description: r.Description,
		Price: values.Price{
			Amount:   uint(r.Price.Amount),
			Currency: r.Price.Currency,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
