package webserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// generateTestData(num, struct) => struct per reflection
func TestGenerateDataForAnyStruct(t *testing.T) {

	t.Run("struct a", func(t *testing.T) {
		expectedAmount := 10
		actual := generateTestData(expectedAmount, TestStructA{})
		expected := []TestStructA{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

		assert.Len(t, expected, expectedAmount)
		assert.Equal(t, expected, actual)
	})
	t.Run("struct b", func(t *testing.T) {
		expectedAmount := 10
		actual := generateTestData(expectedAmount, TestStructB{})
		expected := []TestStructB{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

		assert.Len(t, expected, expectedAmount)
		assert.Equal(t, expected, actual)
	})

}

type TestStructA struct {
}
type TestStructB struct {
}
