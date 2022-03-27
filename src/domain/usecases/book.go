package usecases

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
)

type BookService interface {
	FindOne(
		ctx context.Context,
		fields []string,
		spec dataset.CompositeSpecifier,
	) (*domain.Book, error)

	FindAll(
		ctx context.Context,
		fields []string,
		spec dataset.CompositeSpecifier,
		page dataset.Pager,
		sort dataset.Sorter,
	) (*[]domain.Book, error)

	Count(
		ctx context.Context,
		spec dataset.CompositeSpecifier,
	) (int, error)

	CreateOne(
		ctx context.Context,
		book *domain.Book,
		fields []string,
	) (*domain.Book, error)

	UpdateOne(
		ctx context.Context,
		book *domain.Book,
		fields []string,
		ftu []string,
	) (*domain.Book, error)

	Delete(
		ctx context.Context,
		spec dataset.CompositeSpecifier,
	) (int, error)
}
