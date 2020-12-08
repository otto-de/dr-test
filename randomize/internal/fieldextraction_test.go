package internal

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFindFields(t *testing.T) {

	t.Run("find string field", func(t *testing.T) {
		type StringStruct struct {
			String1 string
			String2 string
		}

		got := getFields(&StringStruct{})
		assert.Len(t, got, 2)
		assert.Equal(t, got[0], Field{
			Name: "String1",
			Kind: reflect.String,
		})
		assert.Equal(t, got[1], Field{
			Name: "String2",
			Kind: reflect.String,
		})

	})

	t.Run("find integer fields", func(t *testing.T) {
		type IntStruct struct {
			Int1 int
			Int2 int8
			Int3 int16
			Int4 int32
			Int5 int64
		}

		got := getFields(&IntStruct{})
		assert.Len(t, got, 5)
		assert.Contains(t, got, Field{Name: "Int1", Kind: reflect.Int})
		assert.Contains(t, got, Field{Name: "Int2", Kind: reflect.Int8})
		assert.Contains(t, got, Field{Name: "Int3", Kind: reflect.Int16})
		assert.Contains(t, got, Field{Name: "Int4", Kind: reflect.Int32})
		assert.Contains(t, got, Field{Name: "Int5", Kind: reflect.Int64})

	})

	t.Run("find float fields", func(t *testing.T) {
		type FloatStruct struct {
			Float1 float32
			Float2 float64
		}

		got := getFields(&FloatStruct{})
		assert.Len(t, got, 2)
		assert.Contains(t, got, Field{Name: "Float1", Kind: reflect.Float32})
		assert.Contains(t, got, Field{Name: "Float2", Kind: reflect.Float64})

	})

	t.Run("find boolean fields", func(t *testing.T) {
		type BooleanStruct struct {
			Boolean bool
		}

		got := getFields(&BooleanStruct{})
		assert.Len(t, got, 1)
		assert.Contains(t, got, Field{Name: "Boolean", Kind: reflect.Bool})
	})

	t.Run("find struct fields", func(t *testing.T) {

		type AnotherStruct struct{}

		type StructStruct struct {
			Struct AnotherStruct
		}

		got := getFields(&StructStruct{})
		assert.Len(t, got, 1)
		assert.Contains(t, got, Field{Name: "Struct", Kind: reflect.Struct})
	})
	t.Run("find slice fields", func(t *testing.T) {
		type StructStruct struct {
			Slice []interface{}
		}

		got := getFields(&StructStruct{})
		assert.Len(t, got, 1)
		assert.Contains(t, got, Field{Name: "Slice", Kind: reflect.Slice})
	})

}
