package internal

import (
	"drtest/randomize/api"
	"fmt"
	"reflect"
)

func Randomize(strukt interface{}, configuration api.Configuration) interface{} {
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

func fillNestedStruct(strukt interface{}, m fieldMeta, configuration api.Configuration) {
	elem := reflect.ValueOf(strukt).Elem()
	field := (&elem).FieldByName(m.Name)
	nestedStruct := reflect.New(field.Type())
	fmt.Printf("Nested struct\n%+v\n", nestedStruct)
	Randomize(nestedStruct.Interface(), configuration)

	field.Set(nestedStruct.Elem())
}

func fillSimpleValue(strukt interface{}, fieldMeta fieldMeta, configuration api.Configuration) {

	elem := reflect.ValueOf(strukt).Elem()
	f := elem.FieldByName(fieldMeta.Name)
	fmt.Printf("Trying value %v\n", fieldMeta)
	if f.CanSet() {
		fmt.Printf("Setting value %v\n", fieldMeta)
		setRandomValue(f, fieldMeta, configuration)
	}
}

func setRandomValue(struktField reflect.Value, fieldMeta fieldMeta, configuration api.Configuration) {
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
		size := randomIntCapped(configuration.MaxListSize)
		slice := randomSlice(sliceType, size)
		fmt.Printf("List size %d", size)
		struktField.Set(slice.Slice(0, size))
	default:
		fmt.Printf("%v not supported", struktField.Kind())
	}
}
