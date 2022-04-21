package repositories

import (
	"context"
	"github.com/aso779/go-ddd-example/infrastructure/connection"
	"github.com/aso779/go-ddd/domain/usecase/dataset"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/dataspec"
	"github.com/uptrace/bun"
	"strings"
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

func (r CrudRepository[E, T]) FindOneByPk(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	pk metadata.PrimaryKey,
) (*E, error) {
	spec := dataspec.NewAnd()
	for k, v := range pk {
		spec.Append(dataspec.NewEqual(k, v))
	}

	return r.FindOne(ctx, tx, fields, spec)
}

func (r CrudRepository[E, T]) FindAll(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	spec dataset.Specifier,
) (*[]E, error) {
	var entities []E
	if tx == nil {
		tx = r.ConnSet.ReadPool()
	}

	query := tx.
		NewSelect().
		Model(&entities).
		Column(fields...)
	if spec != nil && !spec.IsEmpty() {
		for _, j := range spec.Joins(r.Meta) {
			query.Join(j)
		}

		query.Where(spec.Query(r.Meta), spec.Values()...)
	}

	err := query.Scan(ctx)

	return &entities, err
}

func (r CrudRepository[E, T]) FindPage(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	spec dataset.Specifier,
	page dataset.Pager,
	sort dataset.Sorter,
) (*[]E, error) {
	var entities []E
	if tx == nil {
		tx = r.ConnSet.ReadPool()
	}

	query := tx.
		NewSelect().
		Model(&entities).
		Column(fields...)
	if spec != nil && !spec.IsEmpty() {
		for _, j := range spec.Joins(r.Meta) {
			query.Join(j)
		}

		query.Where(spec.Query(r.Meta), spec.Values()...)
	}
	if page != nil && !page.IsEmpty() {
		query.Limit(page.GetSize())
		query.Offset(page.GetOffset())
	}
	if sort != nil && !sort.IsEmpty() {
		query.OrderExpr(sort.OrderBy(r.Meta))
	}

	err := query.Scan(ctx)

	return &entities, err
}

func (r CrudRepository[E, T]) FindAllByPks(
	ctx context.Context,
	tx bun.IDB,
	fields []string,
	pks []metadata.PrimaryKey,
) (*[]E, error) {
	var keys []string
	var values []any
	var isComposite bool
	var spec dataset.Specifier

	for i, pk := range pks {
		if i == 0 {
			isComposite = pk.IsComposite()
			keys = pk.Keys()
		}

		if isComposite {
			var valuesGroup []any
			for _, vv := range pk {
				valuesGroup = append(valuesGroup, vv)
			}
			values = append(values, valuesGroup)
		} else {
			for _, vv := range pk {
				values = append(values, vv)
			}
		}
	}

	if isComposite {
		spec = dataspec.NewCompositeIn(keys, bun.In(values))
	} else {
		spec = dataspec.NewIn(keys[0], bun.In(values))
	}

	return r.FindAll(ctx, tx, fields, spec)
}

func (r CrudRepository[E, T]) Count(
	ctx context.Context,
	tx bun.IDB,
	spec dataset.Specifier,
) (int, error) {
	var ent E
	if tx == nil {
		tx = r.ConnSet.ReadPool()
	}

	query := tx.
		NewSelect().
		Model(&ent)

	if spec != nil && !spec.IsEmpty() {
		for _, j := range spec.Joins(r.Meta) {
			query.Join(j)
		}

		query.Where(spec.Query(r.Meta), spec.Values()...)
	}

	return query.Count(ctx)
}

func (r CrudRepository[E, T]) CreateOne(
	ctx context.Context,
	tx bun.IDB,
	ent *E,
	fields []string,
) (*E, error) {
	if tx == nil {
		tx = r.ConnSet.WritePool()
	}

	_, err := tx.NewInsert().
		Model(ent).
		Returning(strings.Join(fields, ",")).
		Exec(ctx)

	return ent, err
}

func (r CrudRepository[E, T]) UpdateOne(
	ctx context.Context,
	tx bun.IDB,
	ent *E,
	fields []string,
) (*E, error) {
	if tx == nil {
		tx = r.ConnSet.WritePool()
	}

	_, err := tx.NewUpdate().
		Model(ent).
		WherePK().
		Returning(strings.Join(fields, ",")).
		Exec(ctx)

	return ent, err
}

func (r CrudRepository[E, T]) Delete(
	ctx context.Context,
	tx bun.IDB,
	spec dataset.Specifier,
) (int, error) {
	var ent E
	if tx == nil {
		tx = r.ConnSet.WritePool()
	}

	query := tx.NewDelete().
		Model(&ent)
	if spec != nil && !spec.IsEmpty() {
		query.Where(spec.Query(r.Meta), spec.Values()...)
	}

	res, err := query.Exec(ctx)
	if err != nil {
		return 0, err
	}

	rows, err := res.RowsAffected()

	return int(rows), err
}
