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
	fmt.Printf("Created struct %+v", strukt)
	return strukt
}

func setValue(strukt interface{}, fieldMeta FieldMeta) {
	elem := reflect.ValueOf(strukt).Elem()
	f := elem.FieldByName(fieldMeta.Name)
	if f.CanSet() {
		setRandomValue(f, fieldMeta)
	} else {
		fmt.Printf("Cannot set field %q", fieldMeta)
	}
}

func setRandomValue(struktField reflect.Value, fieldMeta FieldMeta) {
	switch fieldMeta.Kind {
	case reflect.String:
		struktField.SetString(randomString())
	}
}

func randomString() string {
	return "foo"
}
