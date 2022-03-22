package meta

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type BookMeta struct {
	domain.Book
}

func (r BookMeta) Entity() metadata.Entity {
	return r.Book
}

func (r BookMeta) Relations() (relations map[string]metadata.Relation) {
	return
}
