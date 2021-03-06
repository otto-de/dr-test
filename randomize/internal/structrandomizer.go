package internal

import (
	"drtest/randomize/api"
	"drtest/randomize/internal/fieldextraction"
	"drtest/randomize/internal/randomContent"
	"fmt"
	"github.com/cjrd/allocate"
	"log"
	"reflect"
)

func Randomize(strukt interface{}, configuration api.Configuration) interface{} {
	allocate.MustZero(strukt)
	fields := fieldextraction.ExtractFieldMetadata(strukt)
	for _, m := range fields {
		fillSimpleValue(strukt, m, configuration)
	}
	return strukt
}

func fillSimpleValue(strukt interface{}, fieldMeta fieldextraction.FieldMeta, configuration api.Configuration) {
	elem := reflect.ValueOf(strukt).Elem()
	f := elem.FieldByName(fieldMeta.Name)
	if f.CanSet() {
		if fieldMeta.IsContainer() {
			f.Set(buildContainer(fieldMeta, configuration))
		} else {
			f.Set(buildValue(fieldMeta, configuration).Convert(fieldMeta.ReflectionType()))
		}
	}
}

func buildContainer(meta fieldextraction.FieldMeta, configuration api.Configuration) reflect.Value {
	if meta.Kind == reflect.Slice {
		return buildRandomSlice(meta, configuration)
	} else {
		return buildRandomMap(meta, configuration)
	}
}

func buildValue(fieldMeta fieldextraction.FieldMeta, configuration api.Configuration) reflect.Value {
	typ := fieldMeta.ReflectionType()
	switch typ.Kind() {
	case reflect.Int, reflect.Uint8, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.String, reflect.Bool:
		return randomContent.RandomSimpleValue(typ, configuration)
	case reflect.Ptr:
		newStruct := reflect.New(fieldMeta.Value.Elem().Type())
		Randomize(newStruct.Interface(), configuration)
		return newStruct

	case reflect.Struct:
		nestedStruct := reflect.New(typ)
		Randomize(nestedStruct.Interface(), configuration)
		return reflect.ValueOf(nestedStruct.Elem().Interface())

	default:
		log.Fatalf("%v not supported.\n", typ.Kind())
		return reflect.Value{}
	}

}

func buildRandomMap(fieldMeta fieldextraction.FieldMeta, configuration api.Configuration) reflect.Value {
	m, size := randomContent.EmptyMap(fieldMeta.ReflectionType(), configuration)
	for i := 0; i < size; i++ {
		key := fmt.Sprintf("%d%v", i, randomContent.RandomString(configuration))
		value := buildValue(fieldMeta, configuration).Interface()
		m.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
	}
	return m

}

func buildRandomSlice(fieldMeta fieldextraction.FieldMeta, configuration api.Configuration) reflect.Value {
	slice, size := randomContent.EmptySlice(fieldMeta.ReflectionType(), configuration)
	for i := 0; i < size; i++ {
		slice = reflect.Append(slice, reflect.ValueOf(buildValue(fieldMeta, configuration).Interface()))
	}
	return slice.Slice(0, size)

}
