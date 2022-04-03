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
}
