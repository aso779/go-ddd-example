package adapters

import (
	"github.com/aso779/go-ddd-example/infrastructure"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
)

type BookSort struct {
	ID        *SortDirection `json:"id"`
	Title     *SortDirection `json:"title"`
	CreatedAt *SortDirection `json:"createdAt"`
}

func (r *BookSort) Build() dataset.Sorter {
	sorter := infrastructure.NewSorter()
	if r != nil {
		if r.ID != nil {
			sorter.Append("id", r.ID.String())
		}

		if r.Title != nil {
			sorter.Append("title", r.Title.String())
		}

		if r.CreatedAt != nil {
			sorter.Append("createdAt", r.CreatedAt.String())
		}

	}

	return sorter
}
