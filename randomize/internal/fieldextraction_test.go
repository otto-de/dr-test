package internal

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFindFields(t *testing.T) {

	assertFieldMeta := func(t *testing.T, data []fieldMeta, index int, name string, kind reflect.Kind) {
		t.Helper()
		assert.True(t, index <= len(data), "index %d out of bounds", index)
		meta := data[index]
		assert.Equal(t, meta.Name, name, "name does not match")
		assert.Equal(t, meta.Kind, kind, "kind does not match")
		assert.NotNil(t, meta.Value, "value is null")
	}

	t.Run("find string field", func(t *testing.T) {
		type StringStruct struct {
			String1 string
			String2 string
		}

		got := getFieldMeta(&StringStruct{})
		assert.Len(t, got, 2)
		assertFieldMeta(t, got, 0, "String1", reflect.String)
		assertFieldMeta(t, got, 1, "String2", reflect.String)

	})

	t.Run("find integer fields", func(t *testing.T) {
		type IntStruct struct {
			Int1 int
			Int2 int8
			Int3 int16
			Int4 int32
			Int5 int64
		}

		got := getFieldMeta(&IntStruct{})
		assert.Len(t, got, 5)
		assertFieldMeta(t, got, 0, "Int1", reflect.Int)
		assertFieldMeta(t, got, 1, "Int2", reflect.Int8)
		assertFieldMeta(t, got, 2, "Int3", reflect.Int16)
		assertFieldMeta(t, got, 3, "Int4", reflect.Int32)
		assertFieldMeta(t, got, 4, "Int5", reflect.Int64)

	})

	t.Run("find float fields", func(t *testing.T) {
		type FloatStruct struct {
			Float1 float32
			Float2 float64
		}

		got := getFieldMeta(&FloatStruct{})
		assert.Len(t, got, 2)
		assertFieldMeta(t, got, 0, "Float1", reflect.Float32)
		assertFieldMeta(t, got, 1, "Float2", reflect.Float64)

	})

	t.Run("find boolean fields", func(t *testing.T) {
		type BooleanStruct struct {
			Boolean bool
		}

		got := getFieldMeta(&BooleanStruct{})
		assert.Len(t, got, 1)
		assertFieldMeta(t, got, 0, "Boolean", reflect.Bool)
	})

	t.Run("find struct fields", func(t *testing.T) {

		type AnotherStruct struct{}

		type StructStruct struct {
			Struct AnotherStruct
		}

		got := getFieldMeta(&StructStruct{})
		assert.Len(t, got, 1)
		assertFieldMeta(t, got, 0, "Struct", reflect.Struct)
	})

	t.Run("find slice fields", func(t *testing.T) {
		type SliceStruct struct {
			Slice []interface{}
		}

		got := getFieldMeta(&SliceStruct{})
		assert.Len(t, got, 1)
		assertFieldMeta(t, got, 0, "Slice", reflect.Slice)
	})

	t.Run("find mixed fields", func(t *testing.T) {
		type MixedStruct struct {
			String string
			Int    int
			Slice  []interface{}
		}

		got := getFieldMeta(&MixedStruct{})
		assert.Len(t, got, 3)
		assertFieldMeta(t, got, 0, "String", reflect.String)
		assertFieldMeta(t, got, 1, "Int", reflect.Int)
		assertFieldMeta(t, got, 2, "Slice", reflect.Slice)
	})

}
