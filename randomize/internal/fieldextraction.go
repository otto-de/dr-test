package internal

import (
	"fmt"
	"reflect"
)

func getFieldMeta(strukt interface{}) []fieldMeta {
	var fields []fieldMeta
	elem := reflect.ValueOf(strukt).Elem()
	elemT := reflect.TypeOf(strukt).Elem()

	for i := 0; i < elem.NumField(); i++ {
		fieldT := elemT.Field(i)
		val := elem.Field(i)
		fields = append(fields, fieldMeta{fieldT.Name, fieldT.Type.Kind(), val})

	}

	return fields
}

type fieldMeta struct {
	Name  string
	Kind  reflect.Kind
	Value reflect.Value
}

func (f fieldMeta) String() string {
	return fmt.Sprintf("%v:%v", f.Name, f.Kind.String())
}
