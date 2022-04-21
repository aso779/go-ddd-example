package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type GenreOneCreateInput struct {
	ParentID int    `json:"parentId"`
	Title    string `json:"title"`
}

func (r *GenreOneCreateInput) ToEntity() *domain.Genre {
	return &domain.Genre{
		ParentID:  r.ParentID,
		Title:     r.Title,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
