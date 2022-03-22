package adapters

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"time"
)

type BookOutput struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBook() *BookOutput {
	return &BookOutput{}
}

func (r *BookOutput) ToOutput(d *domain.Book) *BookOutput {
	res := &BookOutput{
		ID:        d.ID,
		Title:     d.Title,
		CreatedAt: d.CreatedAt.In(time.Local),
		UpdatedAt: d.UpdatedAt.In(time.Local),
	}

	return res
}

func NewBookPage(opts []*BookOutput, page *infrastructure.Page, totalCount int) *BookPage {
	res := &BookPage{Items: opts}
	if page != nil {
		res.PageInfo = &infrastructure.PageInfo{
			Size:       page.GetSize(),
			Number:     page.GetNumber(),
			TotalCount: totalCount,
		}
	}
	return res
}
