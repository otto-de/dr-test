package webserver

import "reflect"

func FieldsAndNames(data interface{}) Fields {
	obj := reflect.New(reflect.TypeOf(data)).Elem()
	typeOf := obj.Type()

	m := make(map[string]reflect.Kind)

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		m[typeOf.Field(i).Name] = field.Type().Kind()

	}
	return Fields{m, []Fields{}}
}

type Fields struct {
	FieldByName map[string]reflect.Kind
	Children    []Fields
}
