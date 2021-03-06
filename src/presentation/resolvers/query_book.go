package resolvers

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"github.com/aso779/go-ddd-example/presentation/adapters"
)

func (r *queryResolver) BookOne(
	ctx context.Context,
	filter adapters.BookFilter,
) (res *adapters.BookOutput, err error) {
	meta := r.metaContainer.Get(domain.Book{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	ent, err := r.services.Book.FindOne(ctx, fields, filter.Build())
	if err != nil {
		return
	}

	res = adapters.NewBook(NewBookRelations(r.metaContainer)).ToOutput(ent)

	return
}

func (r *queryResolver) BookPage(
	ctx context.Context,
	filter *adapters.BookFilter,
	page *infrastructure.Page,
	sort *adapters.BookSort,
) (res *adapters.BookPage, err error) {
	meta := r.metaContainer.Get(domain.Book{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	spec := filter.Build()
	sorter := sort.Build()

	totalCount, err := r.services.Book.Count(ctx, spec)
	if err != nil {
		return
	}

	ents, err := r.services.Book.FindPage(ctx, fields, spec, page, sorter)
	if err != nil {
		return
	}

	opts := make([]*adapters.BookOutput, 0, len(*ents))

	for _, v := range *ents {
		opts = append(opts, adapters.NewBook(NewBookRelations(r.metaContainer)).ToOutput(&v))
	}
	res = adapters.NewBookPage(opts, page, totalCount)

	return
}
