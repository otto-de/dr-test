package webserver

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFieldsAndNames(t *testing.T) {

	t.Run("struct in struct", func(t *testing.T) {
		type Embedded struct {
			Name string
		}

		type NestedStruct struct {
			Foo Embedded
		}

		// when
		names := FieldsAndNames(NestedStruct{})

		// then
		assert.Equal(t, reflect.Struct, names.FieldByName["Foo"])
	})

	t.Run("multiple fields", func(t *testing.T) {
		// given
		type TestStruct struct {
			Name  string
			Int   int64
			Float float64
			Bool  bool
		}

		// when
		names := FieldsAndNames(TestStruct{})

		// then
		assert.Equal(t, reflect.String, names.FieldByName["Name"])
		assert.Equal(t, reflect.Int64, names.FieldByName["Int"])
		assert.Equal(t, reflect.Float64, names.FieldByName["Float"])
		assert.Equal(t, reflect.Bool, names.FieldByName["Bool"])
	})

}
