package services

import (
	"context"
	"database/sql"
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
	authors []int,
) (*domain.Book, error) {
	tx, err := r.bookRepo.ConnSet.WritePool().BeginTx(ctx, &sql.TxOptions{})

	book, err = r.bookRepo.CrudRepository.CreateOne(ctx, tx, book, fields)
	if err != nil {
		r.log.Error("book create", zap.Error(err))
		rollErr := tx.Rollback()
		if rollErr != nil {
			r.log.Error("book create", zap.Error(rollErr))
		}
		return nil, err
	}

	for _, v := range authors {
		err = r.bookRepo.AddAuthor(ctx, tx, book.ID, v)
		if err != nil {
			r.log.Error("add author", zap.Error(err))
			rollErr := tx.Rollback()
			if rollErr != nil {
				r.log.Error("add author", zap.Error(rollErr))
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		r.log.Error("book create", zap.Error(err))
		return nil, err
	}

	return book, nil
}

func (r BookService) UpdateOne(
	ctx context.Context,
	book *domain.Book,
	fields []string,
	authors []int,
) (*domain.Book, error) {
	tx, err := r.bookRepo.ConnSet.WritePool().BeginTx(ctx, &sql.TxOptions{})

	ent, err := r.bookRepo.CrudRepository.FindOneByPk(ctx, tx, []string{"*"}, book.PrimaryKey())
	if err != nil {
		r.log.Error("book update", zap.Error(err))
		rollErr := tx.Rollback()
		if rollErr != nil {
			r.log.Error("book update", zap.Error(rollErr))
		}
		return nil, err
	}

	book.ToExistsEntity(ent)

	book, err = r.bookRepo.CrudRepository.UpdateOne(ctx, tx, ent, fields)

	err = r.bookRepo.DeleteAuthors(ctx, tx, book.ID)
	if err != nil {
		r.log.Error("book update", zap.Error(err))
		rollErr := tx.Rollback()
		if rollErr != nil {
			r.log.Error("book update", zap.Error(rollErr))
		}
		return nil, err
	}

	for _, v := range authors {
		err = r.bookRepo.AddAuthor(ctx, tx, book.ID, v)
		if err != nil {
			r.log.Error("add author", zap.Error(err))
			rollErr := tx.Rollback()
			if rollErr != nil {
				r.log.Error("add author", zap.Error(rollErr))
			}
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		r.log.Error("book update", zap.Error(err))
		return nil, err
	}

	return book, nil
}

func (r BookService) Delete(
	ctx context.Context,
	spec dataset.CompositeSpecifier,
) (int, error) {
	return r.bookRepo.Delete(ctx, nil, spec)
}
