package webserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// generateTestData(num, struct) => struct per reflection
func TestGenerateDataForAnyStruct(t *testing.T) {
	/*t.Run("empty struct", func(t *testing.T) {
		// given
		type EmptyStruct struct{}
		expectedAmount := 10

		// when
		actual := generateTestData(expectedAmount, EmptyStruct{})

		// then
		expected := []EmptyStruct{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

		assert.Len(t, expected, expectedAmount)
		assert.Equal(t, expected, actual)
	})*/

	t.Run("string struct", func(t *testing.T) {
		// given
		type StringStruct struct {
			Name       string
			WasAnderes int64
		}
		expectedAmount := 5

		// when
		actual := generateTestData(expectedAmount, StringStruct{})

		// then
		expected := []StringStruct{{"foo", 0}, {"foo", 0}, {"foo", 0},
			{"foo", 0}, {"foo", 0}}

		assert.Len(t, expected, expectedAmount)
		assert.Equal(t, expected, actual)
	})
}
