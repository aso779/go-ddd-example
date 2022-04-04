package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/domain/projections"
	"time"
)

type AuthorOutput struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewAuthor() *AuthorOutput {
	return &AuthorOutput{}
}

func (r *AuthorOutput) ToOutput(d *domain.Author) *AuthorOutput {
	res := &AuthorOutput{
		ID:        d.ID,
		Name:      d.Name,
		CreatedAt: d.CreatedAt.In(time.Local),
		UpdatedAt: d.UpdatedAt.In(time.Local),
	}

	return res
}

func (r *AuthorOutput) FromProjectionBookId(d *projections.AuthorBookID) *AuthorOutput {
	res := &AuthorOutput{
		ID:        d.ID,
		Name:      d.Name,
		CreatedAt: d.CreatedAt.In(time.Local),
		UpdatedAt: d.UpdatedAt.In(time.Local),
	}

	return res
}
