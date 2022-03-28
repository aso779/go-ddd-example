package infrastructure

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/vektah/gqlparser/v2/ast"
	"strings"
)

func GetPreloads(ctx context.Context, meta metadata.Meta) []string {
	requested := GetNestedPreloads(
		graphql.GetOperationContext(ctx),
		graphql.CollectFieldsCtx(ctx, nil),
		"",
	)
	available := meta.PresenterPersistenceMapping()
	var fields []string
	for _, field := range requested {
		if v, ok := available[field]; ok {
			fields = append(fields, v)
		}
	}
	return fields
}

func GetNestedPreloads(
	ctx *graphql.OperationContext,
	fields []graphql.CollectedField,
	prefix string,
) (preloads []string) {
	for _, column := range fields {
		prefixColumn := GetPreloadString(prefix, column.Name)
		preloads = append(preloads, prefixColumn)
		preloads = append(preloads, GetNestedPreloads(ctx, graphql.CollectFields(ctx, column.Selections, nil), prefixColumn)...)
	}
	return
}

func GetPreloadString(prefix, name string) string {
	if len(prefix) > 0 {
		return prefix + "_" + name
	}
	return name
}

func ParseInputFromContext(ctx context.Context) *ast.Argument {
	arguments := graphql.GetFieldContext(ctx).Field.Arguments

	for _, argument := range arguments {
		if argument.Name == "input" {
			return argument
		}
	}

	return nil
}

func ParseInputChildren(c *ast.Value, prefix ...string) (res []string) {
	for _, field := range c.Children {
		if field.Value.Children != nil {
			res = append(ParseInputChildren(field.Value, prefix...))
		}
		res = append(res, strings.Join(append(prefix, field.Name), "_"))
	}

	return
}

func ParseInputFields(ctx context.Context) (res []string) {
	input := ParseInputFromContext(ctx)

	// for single insert/update
	//if input.Value.Kind != ast.Variable {
	for _, field := range input.Value.Children {
		if field.Value.Children != nil {
			res = append(res, ParseInputChildren(field.Value, field.Name)...)
		} else {
			res = append(res, field.Name)
		}

	}
	//}

	return
}
