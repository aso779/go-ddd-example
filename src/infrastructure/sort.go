package infrastructure

import (
	"fmt"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"strings"
)

func NewSorter() *Sort {
	return &Sort{
		directions: make(map[string]string, 0),
	}
}

type Sort struct {
	directions map[string]string
}

func (r *Sort) OrderBy(meta metadata.Meta) string {
	var clause []string
	for k, v := range r.directions {
		column := meta.PresenterToPersistence(k)
		clause = append(clause, fmt.Sprintf("%s %s", column, strings.ToUpper(v)))
	}

	return strings.Join(clause, ",")
}

func (r *Sort) IsEmpty() bool {
	return len(r.directions) == 0
}

func (r *Sort) Append(field, direction string) {
	r.directions[field] = direction
}
