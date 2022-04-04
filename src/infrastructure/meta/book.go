package meta

import (
	"github.com/aso779/go-ddd-example/domain"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entrel"
)

type BookMeta struct {
	domain.Book
}

func (r BookMeta) Entity() metadata.Entity {
	return r.Book
}

func (r BookMeta) Relations() map[string]metadata.Relation {

	relations := make(map[string]metadata.Relation)

	relations["Genre"] = entrel.ToOne{
		Meta:      Parser(GenreMeta{}),
		JoinTable: "bks_genres",
		JoinColumns: []entrel.JoinColumn{
			{
				Name:           "genre_id",
				ReferencedName: "id",
			},
		},
	}

	relations["Authors"] = entrel.ToMany{
		Meta:      Parser(AuthorMeta{}),
		JoinTable: "bks_authors",
		ViaTable:  "bks_books_authors",
		JoinColumns: []entrel.JoinColumn{
			{
				Name:           "books_id",
				ReferencedName: "bks_books.id",
			},
		},
		InverseJoinColumns: []entrel.JoinColumn{
			{
				Name:           "author_id",
				ReferencedName: "bks_authors.id",
			},
		},
	}

	return relations
}
