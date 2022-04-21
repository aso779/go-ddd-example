package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/domain/values"
	"time"
)

type BookOneUpdateInput struct {
	ID          int        `json:"id"`
	GenreID     int        `json:"genreId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       PriceInput `json:"price"`
	Authors     []int      `json:"authors"`
}

func (r *BookOneUpdateInput) ToEntity() *domain.Book {
	res := &domain.Book{
		ID:          r.ID,
		GenreID:     r.GenreID,
		Title:       r.Title,
		Description: r.Description,
		Price: values.Price{
			Amount:   uint(r.Price.Amount),
			Currency: r.Price.Currency,
		},
		UpdatedAt: time.Now(),
	}

	return res
}
