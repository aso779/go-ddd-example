// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package adapters

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/aso779/go-ddd-example/infrastructure"
)

type BookPage struct {
	Items    []*BookOutput            `json:"items"`
	PageInfo *infrastructure.PageInfo `json:"pageInfo"`
}

type DateFilter struct {
	Eq  *time.Time `json:"eq"`
	Gt  *time.Time `json:"gt"`
	Gte *time.Time `json:"gte"`
	Lt  *time.Time `json:"lt"`
	Lte *time.Time `json:"lte"`
}

type IntFilter struct {
	Eq *int  `json:"eq"`
	In []int `json:"in"`
}

type PriceInput struct {
	// amount
	Amount int `json:"amount"`
	// currency
	Currency string `json:"currency"`
}

type TextFilter struct {
	Search        string `json:"search"`
	CaseSensitive bool   `json:"caseSensitive"`
}

type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
	SortDirectionNone SortDirection = "none"
)

var AllSortDirection = []SortDirection{
	SortDirectionAsc,
	SortDirectionDesc,
	SortDirectionNone,
}

func (e SortDirection) IsValid() bool {
	switch e {
	case SortDirectionAsc, SortDirectionDesc, SortDirectionNone:
		return true
	}
	return false
}

func (e SortDirection) String() string {
	return string(e)
}

func (e *SortDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SortDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SortDirection", str)
	}
	return nil
}

func (e SortDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
