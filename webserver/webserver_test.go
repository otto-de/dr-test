package webserver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// generateTestData(num, struct) => struct per reflection
func TestGenerateTestDataFromStruct(t *testing.T) {
	// given
	// struct
	expectedAmount := 10
	got := generateTestData(expectedAmount, TestData{})
	want := []TestData{{}, {}, {}, {}, {}, {}, {}, {}, {}, {}}

	assert.Len(t, want, 10)
	assert.Equal(t, want, got)

	// when
	// getTestData()

	// then
	// n-Testdaten
}
