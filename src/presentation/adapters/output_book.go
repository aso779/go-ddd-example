package adapters

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"time"
)

type BookOutput struct {
	ID          int       `json:"id"`
	GenreID     int       `json:"genreId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       Price     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	relations BookRelations
}

type Price struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func NewBook(relations BookRelations) *BookOutput {
	return &BookOutput{relations: relations}
}

type BookRelations interface {
	Genre(ctx context.Context, genreId int) (res *GenreOutput, err error)
	Authors(ctx context.Context, bookId int) (res []AuthorOutput, err error)
}

func (r *BookOutput) Genre(ctx context.Context) (res *GenreOutput, err error) {
	return r.relations.Genre(ctx, r.GenreID)
}

func (r *BookOutput) Authors(ctx context.Context) (res []AuthorOutput, err error) {
	return r.relations.Authors(ctx, r.ID)
}

func (r *BookOutput) ToOutput(d *domain.Book) *BookOutput {
	res := &BookOutput{
		ID:          d.ID,
		GenreID:     d.GenreID,
		Title:       d.Title,
		Description: d.Description,
		Price: Price{
			Amount:   int(d.Price.Amount),
			Currency: d.Price.Currency,
		},
		CreatedAt: d.CreatedAt.In(time.Local),
		UpdatedAt: d.UpdatedAt.In(time.Local),

		relations: r.relations,
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
