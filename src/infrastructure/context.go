package infrastructure

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/vektah/gqlparser/v2/ast"
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

func ParseNodesSelectionSet(ctx context.Context, entCache metadata.Meta) (fields []string) {
	rootSet := graphql.CollectFieldsCtx(ctx, nil)
	//TODO optimize
	possibleFields := entCache.PresenterPersistenceMapping()
	for _, column := range rootSet {
		if column.Name == "nodes" {
			nodesCtx := graphql.GetOperationContext(ctx)
			nodesFields := graphql.CollectFields(nodesCtx, column.Selections, nil)
			for _, field := range nodesFields {
				if v, ok := possibleFields[field.Name]; ok {
					fields = append(fields, v)
				}
			}
		}
	}

	return
}

func ParseSelectionFromField(field ast.Field, entCache metadata.Meta) (fields []string) {
	//possibleFields, additionFields := Fields(ent)
	requestedFields := make([]string, len(field.SelectionSet))

	for i := 0; i < len(field.SelectionSet); i++ {
		v := field.SelectionSet[i]
		requestedFields = append(requestedFields, v.(*ast.Field).Name)
	}

	return requestedFields
}

// Parses input from context
func ParseInputFromContext(ctx context.Context) *ast.Argument {
	arguments := graphql.GetFieldContext(ctx).Field.Arguments

	for _, argument := range arguments {
		if argument.Name == "input" {
			return argument
		}
	}

	return nil
}

// Extracts fields from input argument
func ParseInputFields(ctx context.Context, child *string) (res []string) {
	input := ParseInputFromContext(ctx)

	// for single insert/update
	if input.Value.Kind != ast.Variable {
		for _, field := range input.Value.Children {
			if child == nil {
				res = append(res, field.Name)
			} else {
				if field.Name == *child && field.Value.Children != nil { // means nested object
					for _, field2 := range field.Value.Children {
						res = append(res, field2.Name)
					}
				}
			}
		}
	} else {
		vars, ok := graphql.GetOperationContext(ctx).Variables["input"]
		if !ok {
			return
		}
		varsMap := vars.(map[string]interface{})
		for _, field := range input.Value.Definition.Fields {
			if _, ok = varsMap[field.Name]; ok {
				res = append(res, field.Name)
			}
		}
	}

	return
}
