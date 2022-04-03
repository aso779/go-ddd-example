package resolvers

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure"
	"github.com/aso779/go-ddd-example/infrastructure/graph/dataloaders"
	"github.com/aso779/go-ddd-example/presentation/adapters"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
)

type BookRelations struct {
	metaContainer metadata.EntityMetaContainer
}

func NewBookRelations(entsContainer metadata.EntityMetaContainer) adapters.BookRelations {
	return &BookRelations{
		metaContainer: entsContainer,
	}
}

func (r *BookRelations) Genre(
	ctx context.Context,
	genreId int,
) (res *adapters.GenreOutput, err error) {
	meta := r.metaContainer.Get(domain.Genre{}.Name())
	fields := infrastructure.GetPreloads(ctx, meta)
	fields = append(fields, "id")

	return dataloaders.For(ctx).GenreByGenreId.Load(genreId, fields)
}
