package meta

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type AuthorMeta struct {
	domain.Author
}

func (r AuthorMeta) Entity() metadata.Entity {
	return r.Author
}

func (r AuthorMeta) Relations() (relations map[string]metadata.Relation) {
	return
}
