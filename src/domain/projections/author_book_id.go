package projections

import "github.com/aso779/go-ddd-example/domain"

type AuthorBookID struct {
	domain.Author `bun:",extend"`

	BookID int
}
