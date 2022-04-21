package services

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd-example/infrastructure/repositories"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
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
	var pks []metadata.PrimaryKey
	for _, v := range ids {
		pks = append(pks, metadata.PrimaryKey{"id": v})
	}

	return r.genreRepo.CrudRepository.FindAllByPks(ctx, nil, fields, pks)
}

func (r GenreService) CreateOne(
	ctx context.Context,
	genre *domain.Genre,
	fields []string,
) (*domain.Genre, error) {
	return r.genreRepo.CrudRepository.CreateOne(ctx, nil, genre, fields)
}

func (r GenreService) UpdateOne(
	ctx context.Context,
	genre *domain.Genre,
	fields []string,
) (*domain.Genre, error) {
	ent, err := r.genreRepo.CrudRepository.FindOneByPk(ctx, nil, []string{"*"}, genre.PrimaryKey())
	if err != nil {
		return nil, err
	}

	genre.ToExistsEntity(ent)

	return r.genreRepo.CrudRepository.UpdateOne(ctx, nil, ent, fields)
}
