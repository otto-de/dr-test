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
		in := &IntStruct{}
		got := Randomize(in)
		assert.NotNil(t, got)
		intVal := getField(got, "Int").Int()
		assert.True(t, intVal != in.Int, "int has not changed")
	})

	t.Run("randomize float", func(t *testing.T) {
		type FloatStruct struct {
			Float float64
		}
		in := &FloatStruct{}
		got := Randomize(in)
		assert.NotNil(t, got)
		floatVal := getField(got, "Float").Float()
		assert.True(t, floatVal != in.Float, "float has not changed")

	})

	t.Run("randomize boolean", func(t *testing.T) {

	})

}

func getField(strukt interface{}, name string) reflect.Value {
	elem := reflect.ValueOf(strukt).Elem()
	return elem.FieldByName(name)
}
