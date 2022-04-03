package services

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd-example/infrastructure/repositories"
	"go.uber.org/zap"
)

type GenreService struct {
	log       *zap.Logger
	connSet   *connection.ConnSet
	genreRepo *repositories.GenreRepository
}

func NewGenre(
	log *zap.Logger,
	connSet *connection.ConnSet,
	genreRepo *repositories.GenreRepository,
) *GenreService {
	return &GenreService{
		log:       log,
		connSet:   connSet,
		genreRepo: genreRepo,
	}
}

func (r GenreService) FindAllByIds(
	ctx context.Context,
	fields []string,
	ids []int,
) (*[]domain.Genre, error) {
	return r.genreRepo.CrudRepository.FindAllByIds(ctx, nil, fields, ids)
}
