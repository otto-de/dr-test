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
	fmt.Printf("NUMFIELDS %d\n", obj.NumField())

	/*typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
		},
		{
			Name: "Name",
			Type: reflect.TypeOf("abc"),
		},
	})*/

	for i := 0; i < obj.NumField(); i++ {
		field := obj.Field(i)
		switch field.Type().Kind() {
		case reflect.String:
			fmt.Println(field.Type().Name())
			/*fmt.Println(reflect.ValueOf(&obj).CanSet())
			fmt.Println(reflect.ValueOf(reflect.ValueOf(obj).Field(i)).Elem())*/
		}
	}
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
