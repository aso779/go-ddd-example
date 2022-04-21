package resolvers

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"github.com/aso779/go-ddd-example/presentation/adapters"
)

func (r *mutationResolver) AuthorOneCreate(
	ctx context.Context,
	input adapters.AuthorOneCreateInput,
) (res *adapters.AuthorOutput, err error) {
	meta := r.metaContainer.Get(domain.Author{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	ent, err := r.services.Author.CreateOne(ctx, input.ToEntity(), fields)
	if err != nil {
		return
	}

	res = adapters.NewAuthor().ToOutput(ent)

	return
}

func (r *mutationResolver) AuthorOneUpdate(
	ctx context.Context,
	input adapters.AuthorOneUpdateInput,
) (res *adapters.AuthorOutput, err error) {
	meta := r.metaContainer.Get(domain.Author{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	ent, err := r.services.Author.UpdateOne(ctx, input.ToEntity(), fields)
	if err != nil {
		return
	}

	res = adapters.NewAuthor().ToOutput(ent)

	return
}
