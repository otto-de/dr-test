package webserver

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFieldsAndNames(t *testing.T) {
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
	assert.Equal(t, reflect.String, names["Name"])
}
