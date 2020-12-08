package internal

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFindFields(t *testing.T) {
	type StringStruct struct {
		String1 string
		String2 string
	}

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

}
