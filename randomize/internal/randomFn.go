package internal

import (
	"drtest/randomize/api"
	"fmt"
	"math/rand"
	"reflect"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(maxLength int) string {
	b := make([]byte, 1+rand.Intn(maxLength))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomInt63() int32 {
	return rand.Int31()
}
func randomInt64() int64 {
	return rand.Int63()
}

func randomIntCapped(max int) int {
	return rand.Intn(max)
}

func randomFloat() float64 {
	return rand.Float64()
}

func randomBool() bool {
	return rand.Float32() <= 0.5
}

func randomByte() byte {
	if randomBool() {
		return 0
	} else {
		return 1
	}
}

func randomSlice(sliceType reflect.Type, size int) reflect.Value {

	sliceOfType := reflect.SliceOf(sliceType)
	fmt.Printf("Slice type %v", sliceOfType)
	slice := reflect.MakeSlice(sliceOfType, 0, 5)
	fmt.Printf("Created slice %v", slice)
	for i := 0; i < size; i++ {
		slice = reflect.Append(slice, reflect.ValueOf(randomSimpleValue(sliceType)))
	}
	return slice
}

func randomSimpleValue(typ reflect.Type) interface{} {
	switch typ.Kind() {
	case reflect.String:
		return randomString(10)
	case reflect.Int:
		return rand.Int()
	case reflect.Int32, reflect.Int64:
		return randomInt64()
	case reflect.Float32:
		return rand.Float32()
	case reflect.Float64:
		return randomFloat()
	case reflect.Bool:
		return randomBool()
	case reflect.Uint8:
		return randomByte()

	default:

		nestedStruct := reflect.New(typ)
		fmt.Printf("Nested struct\n%+v\n", nestedStruct)
		Randomize(nestedStruct.Interface(), api.Configuration{
			MaxListLength:   3,
			MaxStringLength: 10,
		})

		return nestedStruct.Elem().Interface()

	}
}
