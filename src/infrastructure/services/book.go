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

func (r BookService) FindAll(
	ctx context.Context,
	fields []string,
	spec dataset.CompositeSpecifier,
	page dataset.Pager,
	sort dataset.Sorter,
) (*[]domain.Book, error) {
	return r.bookRepo.CrudRepository.FindAll(ctx, nil, fields, spec, page, sort)
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
	return r.bookRepo.CrudRepository.Create(ctx, nil, book, fields)
}

func (r BookService) UpdateOne(
	ctx context.Context,
	book *domain.Book,
	fields []string,
	ftu []string,
) (*domain.Book, error) {
	//validate
	ent, err := r.bookRepo.CrudRepository.FindOneById(ctx, nil, []string{"*"}, book.PrimaryKey())
	if err != nil {
		return nil, err
	}

	book.ToExistsEntity(ent)
	//validate

	return r.bookRepo.CrudRepository.Update(ctx, nil, ent, fields, ftu)
}

func (r BookService) Delete(
	ctx context.Context,
	spec dataset.CompositeSpecifier,
) (int, error) {
	return r.bookRepo.Delete(ctx, nil, spec)
}
