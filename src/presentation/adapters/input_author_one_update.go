package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type AuthorOneUpdateInput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r *AuthorOneUpdateInput) ToEntity() *domain.Author {
	return &domain.Author{
		ID:        r.ID,
		Name:      r.Name,
		UpdatedAt: time.Now(),
	}
}
