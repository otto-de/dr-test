package internal

import (
	"fmt"
	"reflect"
)

func Randomize(strukt interface{}) interface{} {
	fields := getFieldMeta(strukt)
	for _, m := range fields {
		setValue(strukt, m)
	}
	fmt.Printf("Created struct %+v\n", strukt)
	return strukt
}

func setValue(strukt interface{}, fieldMeta fieldMeta) {
	elem := reflect.ValueOf(strukt).Elem()
	f := elem.FieldByName(fieldMeta.Name)
	if f.CanSet() {
		setRandomValue(f, fieldMeta)
	} else {
		fmt.Printf("Cannot set field %q", fieldMeta)
	}
}

func setRandomValue(struktField reflect.Value, fieldMeta fieldMeta) {
	switch fieldMeta.Kind {
	case reflect.String:
		struktField.SetString(randomString())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		struktField.SetInt(randomInt())
	case reflect.Float32, reflect.Float64:
		struktField.SetFloat(randomFloat())
	case reflect.Bool:
		struktField.SetBool(randomBool())
	case reflect.Slice:
		sliceType := reflect.TypeOf(fieldMeta.Value.Interface()).Elem()
		size := randomIntCapped(100)
		slice := randomSlice(sliceType, size)
		struktField.Set(slice.Slice(0, size))
	}
}
