package resolvers

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"github.com/aso779/go-ddd-example/presentation/adapters"
)

func (r *mutationResolver) BookOneCreate(
	ctx context.Context,
	input adapters.BookOneCreateInput,
) (res *adapters.BookOutput, err error) {
	meta := r.metaContainer.Get(domain.Book{}.Name())
	fields := infrastructure.GetPreloads(ctx, meta)

	ent, err := r.services.Book.CreateOne(ctx, input.ToEntity(), fields)
	if err != nil {
		return
	}

	res = adapters.NewBook(NewBookRelations(r.metaContainer)).ToOutput(ent)

	return
}

func (r *mutationResolver) BookOneUpdate(
	ctx context.Context,
	input adapters.BookOneUpdateInput,
) (res *adapters.BookOutput, err error) {
	meta := r.metaContainer.Get(domain.Book{}.Name())
	fields := infrastructure.GetPreloads(ctx, meta)
	ftu := infrastructure.ParseInputFields(ctx)

	ent, err := r.services.Book.UpdateOne(ctx, input.ToEntity(), fields, meta.PresenterSetToPersistenceSet(ftu))
	if err != nil {
		return
	}

	res = adapters.NewBook(NewBookRelations(r.metaContainer)).ToOutput(ent)

	return
}

func (r *mutationResolver) BookDelete(
	ctx context.Context,
	filter adapters.BookFilter,
) (rows int, err error) {
	rows, err = r.services.Book.Delete(ctx, filter.Build())
	if err != nil {
		return
	}

	return
}
