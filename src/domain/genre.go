package domain

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
	"time"
)

type Genre struct {
	bun.BaseModel `bun:"table:bks_genres,alias:bks_genres"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	ParentID      int       `bun:"parent_id" json:"parentId"`
	Title         string    `bun:"title" json:"title"`
	CreatedAt     time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `bun:"updated_at" json:"updatedAt"`
	DeletedAt     time.Time `bun:",soft_delete,nullzero"`
}

func (r Genre) EntityName() string {
	return "Genre"
}

func (r Genre) PrimaryKey() metadata.PrimaryKey {
	return metadata.PrimaryKey{"id": r.ID}
}

func (r *Genre) ToExistsEntity(exists *Genre) {
	if r.Title != "" {
		exists.Title = r.Title
	}
	exists.UpdatedAt = r.UpdatedAt
}
