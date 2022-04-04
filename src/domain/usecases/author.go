package usecases

import (
	"context"
	"github.com/aso779/go-ddd-example/domain/projections"
)

type AuthorService interface {
	FindAllViaBookIds(
		ctx context.Context,
		fields []string,
		productIds []int,
	) (*[]projections.AuthorBookID, error)
}
