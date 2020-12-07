package webserver

import (
	"fmt"
	"reflect"
)

func FieldsAndNames(data interface{}) Fields {
	obj := reflect.New(reflect.TypeOf(data)).Elem()
	typeOf := obj.Type()
	children := make(map[string]Fields)
	m := make(map[string]reflect.Kind)

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		kind := field.Type().Kind()
		name := typeOf.Field(i).Name
		if kind == reflect.Struct {
			childFields := FieldsAndNames(field.Interface())
			children[name] = childFields
		}

		m[name] = kind

	}

	fmt.Printf("==> %v", children)
	return Fields{m, children}
}

type Fields struct {
	FieldByName map[string]reflect.Kind
	Children    map[string]Fields
}
