package meta

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type GenreMeta struct {
	domain.Genre
}

func (r GenreMeta) Entity() metadata.Entity {
	return r.Genre
}

func (r GenreMeta) Relations() (relations map[string]metadata.Relation) {
	return
}
