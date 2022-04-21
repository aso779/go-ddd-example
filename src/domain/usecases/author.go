package usecases

import (
	"context"
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd-example/domain/projections"
)

type AuthorService interface {
	FindAllViaBookIds(
		ctx context.Context,
		fields []string,
		productIds []int,
	) (*[]projections.AuthorBookID, error)

	CreateOne(
		ctx context.Context,
		book *domain.Author,
		fields []string,
	) (*domain.Author, error)

	UpdateOne(
		ctx context.Context,
		book *domain.Author,
		fields []string,
	) (*domain.Author, error)
}
