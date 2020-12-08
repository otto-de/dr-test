package internal

import (
	"fmt"
	"reflect"
)

func getFieldMeta(strukt interface{}) []fieldMeta {
	fields := []fieldMeta{}
	elem := reflect.ValueOf(strukt).Elem()
	elemT := reflect.TypeOf(strukt).Elem()

	for i := 0; i < elem.NumField(); i++ {
		//field := elem.fieldMeta(i)
		fieldT := elemT.Field(i)
		fields = append(fields, fieldMeta{fieldT.Name, fieldT.Type.Kind()})

	}

	return fields
}

type fieldMeta struct {
	Name string
	Kind reflect.Kind
}

func (f fieldMeta) String() string {
	return fmt.Sprintf("%v:%v", f.Name, f.Kind.String())
}
