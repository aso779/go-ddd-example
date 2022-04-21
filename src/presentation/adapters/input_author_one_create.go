package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type AuthorOneCreateInput struct {
	Name string `json:"name"`
}

func (r *AuthorOneCreateInput) ToEntity() *domain.Author {
	return &domain.Author{
		Name:      r.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
