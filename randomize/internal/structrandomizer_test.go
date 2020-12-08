package internal

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRandomizeSimpleValues(t *testing.T) {

	t.Run("randomize string", func(t *testing.T) {
		type StringStruct struct {
			String string
		}
		got := Randomize(&StringStruct{})
		assert.NotNil(t, got)
		stringField := getField(got, "String").String()
		assert.True(t, len(stringField) > 0, "string is empty")
	})

	t.Run("randomize integer", func(t *testing.T) {
		type IntStruct struct {
			Int int64
		}
		in := &IntStruct{Int: 0}
		got := Randomize(in)
		assert.NotNil(t, got)
		intVal := getField(got, "Int").Int()
		assert.True(t, intVal != 0, "int has not changed")
	})

	t.Run("randomize float", func(t *testing.T) {
		type FloatStruct struct {
			Float float64
		}
		in := &FloatStruct{Float: 0.1}
		got := Randomize(in)
		assert.NotNil(t, got)
		floatVal := getField(got, "Float").Float()
		assert.True(t, floatVal != 0.1, "float has not changed")

	})

	t.Run("randomize boolean", func(t *testing.T) {
		type BooleanStruct struct {
			Boolean bool
		}
		got := Randomize(&BooleanStruct{})
		assert.NotNil(t, got)
	})

	t.Run("randomize multiple simple fields", func(t *testing.T) {
		type MultiStruct struct {
			Boolean bool
			Int32   int32
			Float64 float64
			String  string
		}
		got := Randomize(&MultiStruct{true, 0, 0.0, ""})
		assert.NotNil(t, got)
		assert.True(t, getField(got, "Int32").Int() > 0)
		assert.True(t, getField(got, "Float64").Float() > 0.0)
		assert.True(t, len(getField(got, "String").String()) > 0)
	})

}

func getField(strukt interface{}, name string) reflect.Value {
	elem := reflect.ValueOf(strukt).Elem()
	return elem.FieldByName(name)
}
