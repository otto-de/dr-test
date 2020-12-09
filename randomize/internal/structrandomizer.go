package internal

import (
	"drtest/randomize/api"
	"github.com/cjrd/allocate"
	"log"
	"reflect"
)

func Randomize(strukt interface{}, configuration api.Configuration) interface{} {
	allocate.MustZero(strukt)
	fields := getFieldMeta(strukt)
	log.Printf("Fields to fill: %v\n", fields)
	for _, m := range fields {
		if m.Kind == reflect.Struct {
			fillNestedStruct(strukt, m, configuration)
		} else {
			fillSimpleValue(strukt, m, configuration)
		}
	}
	log.Printf("Created struct\n%+v\n", strukt)
	return strukt
}

func fillNestedStruct(strukt interface{}, m fieldMeta, configuration api.Configuration) {
	elem := reflect.ValueOf(strukt).Elem()
	field := (&elem).FieldByName(m.Name)
	nestedStruct := reflect.New(field.Type())
	Randomize(nestedStruct.Interface(), configuration)

	field.Set(nestedStruct.Elem())
}

func fillSimpleValue(strukt interface{}, fieldMeta fieldMeta, configuration api.Configuration) {

	elem := reflect.ValueOf(strukt).Elem()
	f := elem.FieldByName(fieldMeta.Name)
	if f.CanSet() {
		log.Printf("Setting value %v\n", fieldMeta)
		setRandomValue(f, fieldMeta, configuration)
	}
}

func setRandomValue(fieldToSet reflect.Value, fieldMeta fieldMeta, configuration api.Configuration) {
	switch fieldMeta.Kind {
	case reflect.String:
		fieldToSet.SetString(randomString(configuration.MaxStringLength))
	case reflect.Int32, reflect.Int64:
		fieldToSet.SetInt(randomInt64())
	case reflect.Float32, reflect.Float64:
		fieldToSet.SetFloat(randomFloat())
	case reflect.Bool:
		fieldToSet.SetBool(randomBool())
	case reflect.Slice:
		slice := createRandomSlice(fieldMeta, configuration)
		fieldToSet.Set(slice)
	case reflect.Ptr:
		newStruct := reflect.New(fieldMeta.Value.Elem().Type())
		Randomize(newStruct.Interface(), configuration)
		fieldToSet.Set(newStruct)
	default:
		log.Printf("%v not supported.\n", fieldToSet.Kind())
	}
}

func createRandomSlice(fieldMeta fieldMeta, configuration api.Configuration) reflect.Value {
	sliceType := reflect.TypeOf(fieldMeta.Value.Interface()).Elem()
	size := configuration.MinListLength + randomIntCapped(configuration.MaxListLength)
	if size > configuration.MaxListLength {
		size = configuration.MaxListLength
	}
	slice := randomSlice(sliceType, size)
	return slice.Slice(0, size)

}
