package di

import (
	"github.com/aso779/go-ddd-example/infrastructure/meta"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entmeta"
)

func NewEntities() metadata.EntityMetaContainer {
	c := entmeta.NewContainer()
	c.Add(meta.BookMeta{}, meta.Parser)
	c.Add(meta.GenreMeta{}, meta.Parser)
	c.Add(meta.AuthorMeta{}, meta.Parser)

	return c
}
