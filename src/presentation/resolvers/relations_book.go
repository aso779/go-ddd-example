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
	meta := r.metaContainer.Get(domain.Genre{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)
	fields = append(fields, "id")

	return dataloaders.For(ctx).BookGenreByGenreId.Load(genreId, fields)
}

func (r *BookRelations) Authors(
	ctx context.Context,
	bookId int,
) (res []adapters.AuthorOutput, err error) {
	meta := r.metaContainer.Get(domain.Author{}.EntityName())
	fields := infrastructure.GetPreloads(ctx, meta)

	return dataloaders.For(ctx).BookAuthorsByBookId.Load(bookId, fields)
}
