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

	t.Run("find integer field", func(t *testing.T) {
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

}
