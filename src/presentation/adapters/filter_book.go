package adapters

import (
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/infrastructure/dataspec"
	"github.com/aso779/go-ddd/presentation/filter"
)

type BookFilter struct {
	ID        *IntFilter    `json:"id"`
	Title     *TextFilter   `json:"title"`
	CreatedAt *DateFilter   `json:"createdAt"`
	Genre     *GenreFilter  `json:"genre"`
	Author    *AuthorFilter `json:"author"`
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
		if r.CreatedAt != nil {
			if r.CreatedAt.Eq != nil {
				spec.Append(dataspec.NewEqual(filter.FieldName(parents, "createdAt"), *r.CreatedAt.Eq))
			}
			if r.CreatedAt.Gt != nil {
				spec.Append(dataspec.NewGt(filter.FieldName(parents, "createdAt"), *r.CreatedAt.Gt))
			}
			if r.CreatedAt.Gte != nil {
				spec.Append(dataspec.NewGte(filter.FieldName(parents, "createdAt"), *r.CreatedAt.Gte))
			}
			if r.CreatedAt.Lt != nil {
				spec.Append(dataspec.NewLt(filter.FieldName(parents, "createdAt"), *r.CreatedAt.Lt))
			}
			if r.CreatedAt.Lte != nil {
				spec.Append(dataspec.NewLte(filter.FieldName(parents, "createdAt"), *r.CreatedAt.Lte))
			}
		}
		if r.Genre != nil {
			spec.Append(r.Genre.Build(parents...))
		}
		if r.Author != nil {
			spec.Append(r.Author.Build(parents...))
		}
	}

	return spec
}
