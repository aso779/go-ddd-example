package services

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/domain/projections"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd-example/infrastructure/repositories"
	"go.uber.org/zap"
)

type AuthorService struct {
	log        *zap.Logger
	connSet    *connection.ConnSet
	authorRepo *repositories.AuthorRepository
}

func NewAuthor(
	log *zap.Logger,
	connSet *connection.ConnSet,
	authorRepo *repositories.AuthorRepository,
) *AuthorService {
	return &AuthorService{
		log:        log,
		connSet:    connSet,
		authorRepo: authorRepo,
	}
}

func (r AuthorService) FindAllViaBookIds(
	ctx context.Context,
	fields []string,
	productIds []int,
) (*[]projections.AuthorBookID, error) {
	return r.authorRepo.FindAllViaBookIds(ctx, fields, productIds)
}

func (r AuthorService) CreateOne(
	ctx context.Context,
	author *domain.Author,
	fields []string,
) (*domain.Author, error) {
	return r.authorRepo.CrudRepository.CreateOne(ctx, nil, author, fields)
}

func (r AuthorService) UpdateOne(
	ctx context.Context,
	author *domain.Author,
	fields []string,
) (*domain.Author, error) {
	ent, err := r.authorRepo.CrudRepository.FindOneByPk(ctx, nil, []string{"*"}, author.PrimaryKey())
	if err != nil {
		return nil, err
	}
	author.ToExistsEntity(ent)

	return r.authorRepo.CrudRepository.UpdateOne(ctx, nil, ent, fields)
}
