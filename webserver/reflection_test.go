package webserver

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMakeStruct(t *testing.T) {

	m := make(map[string]reflect.Kind)
	m["Name"] = reflect.String

	fields := Fields{
		FieldByName: m,
		Children:    nil,
	}

	makeStruct := MakeStruct(fields)

	name, found := reflect.TypeOf(makeStruct).FieldByName("Name")
	assert.True(t, found, "found field %s", name)

	name2, found2 := reflect.TypeOf(makeStruct).FieldByName("GibtsNicht")
	assert.False(t, found2, "no field field %s", name2)

}

func TestFieldsAndNames(t *testing.T) {

	t.Run("struct in struct", func(t *testing.T) {

		type Embedded2 struct {
			Number int64
		}

		type Embedded struct {
			Name   string
			Level1 Embedded2
		}

		type NestedStruct struct {
			Level0 Embedded
		}

		// when
		names := FieldsAndNames(NestedStruct{})

		// then

		assert.Equal(t, reflect.Struct, names.FieldByName["Level0"])
		assert.Equal(t, reflect.String, names.Children["Level0"].FieldByName["Name"])

		assert.Equal(t, reflect.Struct, names.Children["Level0"].FieldByName["Level1"])
		assert.Equal(t, reflect.Int64, names.Children["Level0"].Children["Level1"].FieldByName["Number"])
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
