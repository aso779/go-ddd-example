package di

import (
	"fmt"
	"github.com/aso779/go-ddd-example/infrastructure/meta"
	"github.com/aso779/go-ddd/domain/usecase/metadata"
	"github.com/aso779/go-ddd/infrastructure/entmeta"
	"reflect"
	"strings"
)

var parser = func(decorator metadata.EntityMetaDecorator) metadata.Meta {
	m := entmeta.NewMeta()
	m.SetDecorator(decorator)
	m.SetEntityName(decorator.Entity().Name())

	t := reflect.TypeOf(decorator.Entity())

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Name == "BaseModel" {
			tag, _ := t.Field(i).Tag.Lookup("bun")
			tagVals := strings.Split(tag, ",")
			m.SetPersistenceName(strings.TrimPrefix(tagVals[0], "table:"))
			continue
		}

		bunName, bunOk := t.Field(i).Tag.Lookup("bun")
		if !bunOk || strings.Contains(bunName, "fk") || strings.Contains(bunName, "many2many") {
			continue
		}
		bunParts := strings.Split(bunName, ",")

		jsonName, jsonOk := t.Field(i).Tag.Lookup("json")

		if jsonOk {
			m.AddFieldToPresenter(t.Field(i).Name, jsonName)
			m.AddPresenterToPersistence(jsonName, bunParts[0])
			m.AddPersistenceToPresenter(bunParts[0], jsonName)
		}
	}

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
