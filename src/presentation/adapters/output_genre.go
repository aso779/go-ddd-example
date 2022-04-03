package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type GenreOutput struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewGenre() *GenreOutput {
	return &GenreOutput{}
}

func (r *GenreOutput) ToOutput(d *domain.Genre) *GenreOutput {
	res := &GenreOutput{
		ID:        d.ID,
		Title:     d.Title,
		CreatedAt: d.CreatedAt.In(time.Local),
		UpdatedAt: d.UpdatedAt.In(time.Local),
	}

	return res
}
