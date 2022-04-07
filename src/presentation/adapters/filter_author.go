package adapters

import (
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/infrastructure/dataspec"
	"github.com/aso779/go-ddd/presentation/filter"
)

type AuthorFilter struct {
	ID   *IntFilter  `json:"id"`
	Name *TextFilter `json:"name"`
}

func (r *AuthorFilter) Build(parents ...string) dataset.CompositeSpecifier {
	parents = append(parents, "Author")
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
		if r.Name != nil {
			if r.Name.CaseSensitive {
				spec.Append(dataspec.NewLike(filter.FieldName(parents, "name"), r.Name.Search))
			} else {
				spec.Append(dataspec.NewILike(filter.FieldName(parents, "name"), r.Name.Search))
			}
		}
	}

	return spec
}
