package adapters

import (
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/infrastructure/dataspec"
	"github.com/aso779/go-ddd/presentation/filter"
)

type BookFilter struct {
	ID    *IntFilter  `json:"id"`
	Title *TextFilter `json:"title"`
}

func (r *BookFilter) Build(parents ...string) dataset.CompositeSpecifier {
	parents = append(parents, "Book")
	spec := dataspec.NewAnd()

	if r != nil {
		if r.ID != nil {
			if r.ID.Eq != nil {
				spec.Append(dataspec.NewEqual(filter.FieldName(parents, "id"), *r.ID.Eq))
			}
			if r.ID.In != nil {
				spec.Append(dataspec.NewIn(filter.FieldName(parents, "id"), r.ID.In))
			}
		}
		if r.Title != nil {
			if r.Title.CaseSensitive {
				spec.Append(dataspec.NewLike(filter.FieldName(parents, "title"), r.Title.Search))
			} else {
				spec.Append(dataspec.NewILike(filter.FieldName(parents, "title"), r.Title.Search))
			}
		}
	}

	return spec
}
