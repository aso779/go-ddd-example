package resolvers

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"github.com/aso779/go-ddd-example/presentation/adapters"
)

func (r *mutationResolver) GenreOneCreate(
	ctx context.Context,
	input adapters.GenreOneCreateInput,
) (res *adapters.GenreOutput, err error) {
	meta := r.metaContainer.Get(domain.Genre{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	ent, err := r.services.Genre.CreateOne(ctx, input.ToEntity(), fields)
	if err != nil {
		return
	}

	res = adapters.NewGenre().ToOutput(ent)

	return
}

func (r *mutationResolver) GenreOneUpdate(
	ctx context.Context,
	input adapters.GenreOneUpdateInput,
) (res *adapters.GenreOutput, err error) {
	meta := r.metaContainer.Get(domain.Genre{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	ent, err := r.services.Genre.UpdateOne(ctx, input.ToEntity(), fields)
	if err != nil {
		return
	}

	res = adapters.NewGenre().ToOutput(ent)

	return
}
