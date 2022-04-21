package domain

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
	"time"
)

type Author struct {
	bun.BaseModel `bun:"table:bks_authors,alias:bks_authors"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	Name          string    `bun:"name" json:"name"`
	CreatedAt     time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `bun:"updated_at" json:"updatedAt"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

func (r Author) EntityName() string {
	return "Author"
}

func (r Author) PrimaryKey() metadata.PrimaryKey {
	return metadata.PrimaryKey{"id": r.ID}
}

func (r *Author) ToExistsEntity(exists *Author) {
	if r.Name != "" {
		exists.Name = r.Name
	}
	exists.UpdatedAt = r.UpdatedAt
}
