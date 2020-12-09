package internal

import (
	"drtest/randomize/api"
	"fmt"
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
		fieldToSet.Set(buildRandomSlice(fieldMeta, configuration))
	case reflect.Map:
		fieldToSet.Set(buildRandomMap(fieldMeta, configuration))

	case reflect.Ptr:
		newStruct := reflect.New(fieldMeta.Value.Elem().Type())
		Randomize(newStruct.Interface(), configuration)
		fieldToSet.Set(newStruct)
	default:
		log.Printf("%v not supported.\n", fieldToSet.Kind())
	}
}

func buildRandomMap(fieldMeta fieldMeta, configuration api.Configuration) reflect.Value {

	valueType := reflect.TypeOf(fieldMeta.Value.Interface()).Elem()
	mapType := reflect.MapOf(reflect.TypeOf("str"), valueType)
	m := reflect.MakeMap(mapType)

	size := getLen(configuration.MinMapLength, configuration.MaxMapLength)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("%d%v", i, randomString(configuration.MaxStringLength))
		value := randomSimpleValue(valueType)
		fmt.Printf("Key %v Val %v", key, value)
		m.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	}
	return m

}

func buildRandomSlice(fieldMeta fieldMeta, configuration api.Configuration) reflect.Value {
	sliceType := reflect.TypeOf(fieldMeta.Value.Interface()).Elem()
	size := getLen(configuration.MinListLength, configuration.MaxListLength)
	slice := randomSlice(sliceType, size)
	return slice.Slice(0, size)

}

func getLen(minLen, maxLen int) int {
	if maxLen <= 0 {
		return 1
	}
	proposed := randomIntCapped(maxLen)
	if proposed < minLen {
		return minLen
	}
	return proposed
}
