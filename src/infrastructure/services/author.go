package services

import (
	"context"
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
