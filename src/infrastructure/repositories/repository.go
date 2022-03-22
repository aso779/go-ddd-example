package repositories

import (
	"context"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/uptrace/bun"
)

type CrudRepository[E metadata.Entity, T bun.Tx] struct {
	ConnSet *connection.ConnSet
	Meta    metadata.Meta
}

func (r CrudRepository[E, T]) FindOne(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	spec dataset.Specifier,
) (*E, error) {
	var ent E
	if tx == nil {
		tx = r.ConnSet.ReadPool()
	}

	query := tx.
		NewSelect().
		Model(&ent).
		Column(fields...)

	if spec != nil && !spec.IsEmpty() {
		for _, j := range spec.Joins(r.Meta) {
			query.Join(j)
		}
		query.Where(spec.Query(r.Meta), spec.Values()...)
	}

	err := query.Scan(ctx)

	return &ent, err
}

func (r CrudRepository[E, T]) FindOneById(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	id metadata.PrimaryKey,
) (*E, error) {
	return nil, nil
}

func (r CrudRepository[E, T]) FindAll(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	spec dataset.Specifier,
	page dataset.Pager,
	sort dataset.Sorter,
) (*[]E, error) {
	var ents []E
	if tx == nil {
		tx = r.ConnSet.ReadPool()
	}

	//fields = r.repo.PrepareColumnNames(fields, r.entCache.GetPersistenceName())

	query := tx.
		NewSelect().
		Model(&ents).
		Column(fields...)
	if spec != nil && !spec.IsEmpty() {
		for _, j := range spec.Joins(r.Meta) {
			query.Join(j)
		}

		query.Where(spec.Query(r.Meta), spec.Values()...)
	}
	if page != nil && !page.IsEmpty() {
		query.Limit(page.GetNumber())
		query.Offset(page.GetOffset())
	}
	if sort != nil && !sort.IsEmpty() {
		query.OrderExpr(sort.OrderBy(r.Meta))
	}

	err := query.Scan(ctx)

	return &ents, err
}

func (r CrudRepository[E, T]) FindAllByIds(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	ids []metadata.PrimaryKey,
) (*[]E, error) {
	return nil, nil
}

func (r CrudRepository[E, T]) Count(
	ctx context.Context,
	tx bun.IDB,
	spec dataset.Specifier,
) (int, error) {
	return 0, nil
}

func (r CrudRepository[E, T]) Create(
	ctx context.Context,
	tx bun.IDB,
	ent *E,
	fields []string,
) (*E, error) {
	return nil, nil
}

func (r CrudRepository[E, T]) Update(
	ctx context.Context,
	tx bun.IDB,
	ent *E,
	fields []string,
	ftu []string,
) (*E, error) {
	return nil, nil
}

func (r CrudRepository[E, T]) Delete(
	ctx context.Context,
	tx bun.IDB,
	spec dataset.Specifier,
) (int, error) {
	return 0, nil
}
