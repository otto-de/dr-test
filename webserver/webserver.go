package webserver

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Test struct {
	name string
	age  int
}

func testt() {

	foo := Test{"ima", 2}

	fmt.Println(string(reflect.ValueOf(foo).Field(0).String()))

	field := reflect.ValueOf(&foo).Elem().FieldByName("name")

	reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).
		Elem().
		Set(reflect.ValueOf("new Value"))

	fmt.Println(string(reflect.ValueOf(foo).Field(0).String()))
}

type Foo struct{ Name string }

func instantiate(dataType reflect.Type) interface{} {
	obj := reflect.New(dataType).Elem()

	return obj.Interface()
}

func generateTestData(amount int, data interface{}) interface{} {
	dataType := reflect.TypeOf(data)
	var slice = make([]interface{}, amount)
	//slice[0] = reflect.New(dataType).Interface()
	for i, _ := range slice {
		slice[i] = instantiate(dataType)
	}

	return slice
}
