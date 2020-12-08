package internal

import (
	"fmt"
	"reflect"
)

func Randomize(strukt interface{}, configuration Configuration) interface{} {
	fields := getFieldMeta(strukt)
	fmt.Printf("Fields to fill: %v\n", fields)
	for _, m := range fields {
		if m.Kind == reflect.Struct {
			fillNestedStruct(strukt, m, configuration)
		} else {
			fillSimpleValue(strukt, m, configuration)
		}
	}
	fmt.Printf("Created struct\n%+v\n", strukt)
	return strukt
}

type Configuration struct {
	MaxListSize     int
	MaxStringLength int
}

func fillNestedStruct(strukt interface{}, m fieldMeta, configuration Configuration) {
	elem := reflect.ValueOf(strukt).Elem()
	field := (&elem).FieldByName(m.Name)
	nestedStruct := reflect.New(field.Type())
	Randomize(nestedStruct.Interface(), configuration)
	fmt.Printf("Nested struct\n%+v\n", nestedStruct)
	field.Set(nestedStruct.Elem())
}

func fillSimpleValue(strukt interface{}, fieldMeta fieldMeta, configuration Configuration) {
	elem := reflect.ValueOf(strukt).Elem()
	f := elem.FieldByName(fieldMeta.Name)

	if f.CanSet() {
		setRandomValue(f, fieldMeta, configuration)
	}
}

func setRandomValue(struktField reflect.Value, fieldMeta fieldMeta, configuration Configuration) {
	switch fieldMeta.Kind {
	case reflect.String:
		struktField.SetString(randomString(configuration.MaxStringLength))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		struktField.SetInt(randomInt())
	case reflect.Float32, reflect.Float64:
		struktField.SetFloat(randomFloat())
	case reflect.Bool:
		struktField.SetBool(randomBool())
	case reflect.Slice:
		sliceType := reflect.TypeOf(fieldMeta.Value.Interface()).Elem()
		size := 1 + randomIntCapped(configuration.MaxListSize)
		slice := randomSlice(sliceType, size)
		struktField.Set(slice.Slice(0, size))

	}
}
