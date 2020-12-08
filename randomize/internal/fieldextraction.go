package internal

import (
	"reflect"
)

func getFields(strukt interface{}) []Field {
	fields := []Field{}
	elem := reflect.ValueOf(strukt).Elem()
	elemT := reflect.TypeOf(strukt).Elem()

	for i := 0; i < elem.NumField(); i++ {
		//field := elem.Field(i)
		fieldT := elemT.Field(i)
		fields = append(fields, Field{fieldT.Name, fieldT.Type.Kind()})

	}

	return fields
}

type Field struct {
	Name string
	Kind reflect.Kind
}
