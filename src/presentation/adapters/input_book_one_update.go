package adapters

import "github.com/aso779/go-ddd-example/domain"

type BookOneUpdateInput struct {
	ID    int     `json:"id"`
	Title *string `json:"title"`
}

func (r *BookOneUpdateInput) ToEntity() *domain.Book {
	res := &domain.Book{
		ID:    int64(r.ID),
		Title: *r.Title,
	}

	return res
}
