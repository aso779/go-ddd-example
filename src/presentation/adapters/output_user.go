package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"time"
)

type UserOutput struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser() *UserOutput {
	return &UserOutput{}
}

func (r *UserOutput) ToOutput(d *domain.User) *UserOutput {
	res := &UserOutput{
		ID:        d.ID,
		Email:     d.Email,
		CreatedAt: d.CreatedAt.In(time.Local),
		UpdatedAt: d.UpdatedAt.In(time.Local),
	}

	return res
}
