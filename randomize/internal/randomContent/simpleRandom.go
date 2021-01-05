package randomContent

import (
	"drtest/randomize/api"
	"log"
	"math/rand"
	"reflect"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(configuration api.Configuration) string {
	b := make([]byte, getLen(configuration.MinStringLength, configuration.MaxStringLength))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomBool() bool {
	return rand.Float32() <= 0.5
}

func randomByte() byte {
	rand.Seed(time.Now().UnixNano())
	max := 255
	val := rand.Intn(max + 1)
	return uint8(val)
}

func RandomSimpleValue(typ reflect.Type, configuration api.Configuration) reflect.Value {
	return reflect.ValueOf(randomSimpleValue(typ, configuration))
}

func randomSimpleValue(typ reflect.Type, configuration api.Configuration) interface{} {
	switch typ.Kind() {
	case reflect.String:
		return RandomString(configuration)
	case reflect.Int:
		return rand.Int()
	case reflect.Int32:
		return rand.Int31()
	case reflect.Int64:
		return rand.Int63()
	case reflect.Float32:
		return rand.Float32()
	case reflect.Float64:
		return rand.Float64()
	case reflect.Bool:
		return randomBool()
	case reflect.Uint8:
		return randomByte()
	default:
		log.Printf("Unsupported type %v\n", typ.Kind())
		return nil
	}
}
