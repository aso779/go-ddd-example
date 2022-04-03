package domain

import (
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
	"time"
)

type Genre struct {
	bun.BaseModel `bun:"table:bks_genres,alias:bks_genres"`
	ID            int       `bun:"id,pk,autoincrement" json:"id"`
	Title         string    `bun:"title" json:"title"`
	CreatedAt     time.Time `bun:"created_at" json:"createdAt"`
	UpdatedAt     time.Time `bun:"updated_at" json:"updatedAt"`
}

func (r Genre) Name() string {
	return "Genre"
}

func (r Genre) PrimaryKey() metadata.PrimaryKey {
	return metadata.PrimaryKey{"id": r.ID}
}
