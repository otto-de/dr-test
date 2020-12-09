package fieldextraction

import (
	"fmt"
	"reflect"
)

func ExtractFieldMetadata(strukt interface{}) []FieldMeta {
	var fields []FieldMeta
	elem := reflect.ValueOf(strukt).Elem()
	elemT := reflect.TypeOf(strukt).Elem()

	for i := 0; i < elem.NumField(); i++ {
		fieldT := elemT.Field(i)
		val := elem.Field(i)
		fields = append(fields, FieldMeta{fieldT.Name, fieldT.Type.Kind(), val})

	}

	return fields
}

type FieldMeta struct {
	Name  string
	Kind  reflect.Kind
	Value reflect.Value
}

func (f FieldMeta) IsContainer() bool {
	return f.Value.Kind() == reflect.Slice || f.Value.Kind() == reflect.Map
}

func (f FieldMeta) ReflectionType() reflect.Type {
	switch f.Value.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return reflect.TypeOf(f.Value.Interface()).Elem()
	case reflect.Ptr:
		return f.Value.Type()
	default:
		return reflect.TypeOf(f.Value.Interface())
	}
}

func (f FieldMeta) String() string {
	return fmt.Sprintf("%v:%v", f.Name, f.Kind.String())
}
