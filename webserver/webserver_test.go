package webserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// generateTestData(num, struct) => struct per reflection
func TestGenerateTestDataFromStruct(t *testing.T) {
	expectedAmount := 10
	actual := generateTestData(expectedAmount, TestData{})
	expected := []TestData{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

	assert.Len(t, expected, expectedAmount)
	assert.Equal(t, expected, actual)

}
