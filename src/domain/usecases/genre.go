package usecases

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
)

type GenreService interface {
	FindAllByIds(
		ctx context.Context,
		fields []string,
		ids []int,
	) (*[]domain.Genre, error)

	CreateOne(
		ctx context.Context,
		book *domain.Genre,
		fields []string,
	) (*domain.Genre, error)

	UpdateOne(
		ctx context.Context,
		book *domain.Genre,
		fields []string,
	) (*domain.Genre, error)
}
