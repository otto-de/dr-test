package internal

import (
	"fmt"
	"reflect"
)

func getFieldMeta(strukt interface{}) []FieldMeta {
	fields := []FieldMeta{}
	elem := reflect.ValueOf(strukt).Elem()
	elemT := reflect.TypeOf(strukt).Elem()

	for i := 0; i < elem.NumField(); i++ {
		//field := elem.FieldMeta(i)
		fieldT := elemT.Field(i)
		fields = append(fields, FieldMeta{fieldT.Name, fieldT.Type.Kind()})

	}

	return fields
}

type FieldMeta struct {
	Name string
	Kind reflect.Kind
}

func (f FieldMeta) String() string {
	return fmt.Sprintf("%v:%v", f.Name, f.Kind.String())
}
