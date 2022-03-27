package adapters

import "github.com/aso779/go-ddd-example/domain"

type BookOneCreateInput struct {
	Title string `json:"title"`
}

func (r *BookOneCreateInput) ToEntity() *domain.Book {
	return &domain.Book{
		Title: r.Title,
	}
}
