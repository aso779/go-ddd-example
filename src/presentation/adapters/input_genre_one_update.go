package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type GenreOneUpdateInput struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parentId"`
	Title    string `json:"title"`
}

func (r *GenreOneUpdateInput) ToEntity() *domain.Genre {
	return &domain.Genre{
		ID:        r.ID,
		ParentID:  r.ParentID,
		Title:     r.Title,
		UpdatedAt: time.Now(),
	}
}
