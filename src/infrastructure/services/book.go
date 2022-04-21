package services

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd-example/infrastructure/repositories"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"go.uber.org/zap"
)

type BookService struct {
	log      *zap.Logger
	connSet  *connection.ConnSet
	bookRepo *repositories.BookRepository
}

func NewBook(
	log *zap.Logger,
	connSet *connection.ConnSet,
	bookRepo *repositories.BookRepository,
) *BookService {
	return &BookService{
		log:      log,
		connSet:  connSet,
		bookRepo: bookRepo,
	}
}

func (r BookService) FindOne(
	ctx context.Context,
	fields []string,
	spec dataset.CompositeSpecifier,
) (*domain.Book, error) {
	return r.bookRepo.CrudRepository.FindOne(ctx, nil, fields, spec)
}

func (r BookService) FindPage(
	ctx context.Context,
	fields []string,
	spec dataset.CompositeSpecifier,
	page dataset.Pager,
	sort dataset.Sorter,
) (*[]domain.Book, error) {
	return r.bookRepo.CrudRepository.FindPage(ctx, nil, fields, spec, page, sort)
}

func (r BookService) Count(
	ctx context.Context,
	spec dataset.CompositeSpecifier,
) (int, error) {
	return r.bookRepo.CrudRepository.Count(ctx, nil, spec)
}

func (r BookService) CreateOne(
	ctx context.Context,
	book *domain.Book,
	fields []string,
) (*domain.Book, error) {
	return r.bookRepo.CrudRepository.CreateOne(ctx, nil, book, fields)
}

func (r BookService) UpdateOne(
	ctx context.Context,
	book *domain.Book,
	fields []string,
) (*domain.Book, error) {
	ent, err := r.bookRepo.CrudRepository.FindOneByPk(ctx, nil, []string{"*"}, book.PrimaryKey())
	if err != nil {
		return nil, err
	}

	book.ToExistsEntity(ent)

	return r.bookRepo.CrudRepository.UpdateOne(ctx, nil, ent, fields)
}

func (r BookService) Delete(
	ctx context.Context,
	spec dataset.CompositeSpecifier,
) (int, error) {
	return r.bookRepo.Delete(ctx, nil, spec)
}
