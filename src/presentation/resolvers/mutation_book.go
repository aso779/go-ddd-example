package resolvers

import (
	"context"
	"github.com/aso779/go-ddd-example/presentation/adapters"
)

func (r *mutationResolver) BookDelete(
	ctx context.Context,
	filter adapters.BookFilter,
) (rows int, err error) {

	return
}
