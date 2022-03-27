package di

import (
	"fmt"
	"github.com/aso779/go-ddd-example/infrastructure/meta"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entmeta"
	"reflect"
	"strings"
)

type FieldTags struct {
	Name        string
	Presenter   string
	Persistence string
}

func structParser(t reflect.Type, m *entmeta.Meta, prefix ...string) {
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "BaseModel" {
			tag := t.Field(i).Tag.Get("bun")
			tagVals := strings.Split(tag, ",")
			m.SetPersistenceName(strings.TrimPrefix(tagVals[0], "table:"))
			continue
		}

		tag := t.Field(i).Tag.Get("bun")

		if strings.HasPrefix(tag, "embed:") {
			if t.Field(i).Type.Kind() == reflect.Struct {
				structParser(t.Field(i).Type, m, strings.TrimPrefix(tag, "embed:"))
			}
		}

		if ok, fieldTags := fieldParser(t.Field(i)); ok {
			present := strings.Join(append(prefix, fieldTags.Presenter), "")
			persist := strings.Join(append(prefix, fieldTags.Persistence), "")
			m.AddFieldToPresenter(fieldTags.Name, fieldTags.Presenter)
			m.AddPresenterToPersistence(present, persist)
			m.AddPersistenceToPresenter(persist, present)
		}
	}
}

func fieldParser(field reflect.StructField) (bool, *FieldTags) {
	bunName, bunOk := field.Tag.Lookup("bun")
	if !bunOk || strings.Contains(bunName, "fk") || strings.Contains(bunName, "many2many") {
		return false, nil
	}
	bunParts := strings.Split(bunName, ",")

	jsonName, jsonOk := field.Tag.Lookup("json")

	if jsonOk {
		return true, &FieldTags{
			Name:        field.Name,
			Presenter:   jsonName,
			Persistence: bunParts[0],
		}
	}

	return false, nil
}

var parser = func(decorator metadata.EntityMetaDecorator) metadata.Meta {
	m := entmeta.NewMeta()
	m.SetDecorator(decorator)
	m.SetEntityName(decorator.Entity().Name())

	t := reflect.TypeOf(decorator.Entity())

	structParser(t, m)

	//TODO naming & recursion

	relations := make(map[string]metadata.Relation)
	for k, v := range decorator.Relations() {
		relations[k] = v
		for kk, vv := range v.Meta().Relations() {
			compositeKey := fmt.Sprintf("%s.%s", k, kk)
			relations[compositeKey] = vv
		}
	}

	m.SetRelations(relations)

	return m
}

func NewEntities() metadata.EntityMetaContainer {
	c := entmeta.NewContainer()
	c.Add(meta.BookMeta{}, parser)

	return c
}
