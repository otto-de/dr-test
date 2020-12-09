package internal

import (
	"drtest/randomize/api"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRandomizeSimpleValues(t *testing.T) {

	t.Run("randomize string", func(t *testing.T) {
		type StringStruct struct {
			String string
		}
		got := randomizeWithDefaults(&StringStruct{})
		assert.NotNil(t, got)
		stringField := getField(got, "String").String()
		assert.True(t, len(stringField) > 0, "string is empty")
	})

	t.Run("randomize integer", func(t *testing.T) {
		type IntStruct struct {
			Int int64
		}
		in := &IntStruct{Int: 0}
		got := randomizeWithDefaults(in)
		assert.NotNil(t, got)
		intVal := getField(got, "Int").Int()
		assert.True(t, intVal != 0, "int has not changed")
	})

	t.Run("randomize float", func(t *testing.T) {
		type FloatStruct struct {
			Float float64
		}
		in := &FloatStruct{Float: 0.1}
		got := randomizeWithDefaults(in)
		assert.NotNil(t, got)
		floatVal := getField(got, "Float").Float()
		assert.True(t, floatVal != 0.1, "float has not changed")

	})

	t.Run("randomize boolean", func(t *testing.T) {
		type BooleanStruct struct {
			Boolean bool
		}
		got := randomizeWithDefaults(&BooleanStruct{})
		assert.NotNil(t, got)
	})

	t.Run("randomize multiple simple fields", func(t *testing.T) {
		type MultiStruct struct {
			Boolean bool
			Int32   int32
			Float64 float64
			String  string
		}
		got := randomizeWithDefaults(&MultiStruct{})
		assert.NotNil(t, got)
		assert.True(t, getField(got, "Int32").Int() != 0)
		assert.True(t, getField(got, "Float64").Float() != 0.0)
		assert.True(t, len(getField(got, "String").String()) > 0)
	})

}

func TestRandomizeMaps(t *testing.T) {

	t.Run("map with primitive values", func(t *testing.T) {
		type MapStruct struct {
			Int     map[string]int
			Float32 map[string]float32
			Float64 map[string]float64
			String  map[string]string
			Bool    map[string]bool
		}

		got := randomizeWithDefaults(&MapStruct{}).(*MapStruct)
		assert.NotNil(t, got)
		assert.NotEmpty(t, got.Int)
		assert.NotEmpty(t, got.Float32)
		assert.NotEmpty(t, got.Float64)
		assert.NotEmpty(t, got.String)
		assert.NotEmpty(t, got.Bool)
	})

	t.Run("map with struct values", func(t *testing.T) {
		type StructA struct {
			Name string
			Bool bool
		}

		type StructB struct {
			Map map[string]StructA
		}

		got := randomizeWithDefaults(&StructB{}).(*StructB)
		assert.NotNil(t, got)
		assert.NotEmpty(t, got.Map)
		for k, v := range got.Map {
			assert.NotNil(t, k)
			assert.NotNil(t, v)
			assert.True(t, len(v.Name) > 0)
		}
	})

}

func TestRandomizeSlices(t *testing.T) {
	type StructWithSlice struct {
		String       string
		Slice        []string
		SliceInt64   []int64
		SliceFloat64 []float64
		SliceBool    []bool
	}

	got := randomizeWithDefaults(&StructWithSlice{})
	assert.NotNil(t, got)
	assert.True(t, len(getField(got, "String").String()) > 0)
	assert.NotEmpty(t, getField(got, "Slice").Slice(0, 1))
	assert.NotEmpty(t, getField(got, "SliceInt64").Slice(0, 1))
	assert.NotEmpty(t, getField(got, "SliceFloat64").Slice(0, 1))
	assert.NotEmpty(t, getField(got, "SliceBool").Slice(0, 1))
}

func TestNestedStructs(t *testing.T) {

	type InnerStruct struct {
		Int       int64
		InnerName string
	}

	type MiddleStruct struct {
		MiddleName  string
		InnerStruct InnerStruct
	}

	type OuterStruct struct {
		OuterName    string
		MiddleStruct MiddleStruct
	}

	got := randomizeWithDefaults(&OuterStruct{})
	assert.NotNil(t, got)
	assert.True(t, len(getField(got, "OuterName").String()) > 0)
	assert.True(t, len(getField(got, "MiddleStruct").FieldByName("MiddleName").String()) > 0)
	assert.True(t, len(getField(got, "MiddleStruct").FieldByName("InnerStruct").FieldByName("InnerName").String()) > 0)

}

func TestWithPointer(t *testing.T) {

	type Inner struct {
		Name string
	}

	type PointerStruct struct {
		Name     string
		Pointer2 *Inner
	}

	t.Run("pointer to existing struct", func(t *testing.T) {
		got := randomizeWithDefaults(&PointerStruct{"foo", &Inner{}})
		assert.NotNil(t, got)
		assert.True(t, len(getField(got, "Name").String()) > 0)
		assert.False(t, getField(got, "Pointer2").IsNil())
	})

	t.Run("pointer to nil", func(t *testing.T) {
		got := randomizeWithDefaults(&PointerStruct{})
		assert.NotNil(t, got)
		assert.True(t, len(getField(got, "Name").String()) > 0)
		assert.False(t, getField(got, "Pointer2").IsNil())

		p := got.(*PointerStruct)
		fmt.Printf("Name %q\n", p.Name)
		fmt.Printf("Inner %+v", p.Pointer2)
	})

}

func getField(strukt interface{}, name string) reflect.Value {
	elem := reflect.ValueOf(strukt).Elem()
	return elem.FieldByName(name)
}

func randomizeWithDefaults(strukt interface{}) interface{} {
	return Randomize(strukt, api.DefaultConfiguration())
}
