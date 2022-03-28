package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type BookOneUpdateInput struct {
	ID          int        `json:"id"`
	GenreID     int64      `json:"genreId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Price       PriceInput `json:"price"`
}

func (r *BookOneUpdateInput) ToEntity() *domain.Book {
	res := &domain.Book{
		ID:          int64(r.ID),
		GenreID:     r.GenreID,
		Title:       r.Title,
		Description: r.Description,
		Price: domain.Price{
			Amount:   uint64(r.Price.Amount),
			Currency: r.Price.Currency,
		},
		UpdatedAt: time.Now(),
	}

	return res
}
