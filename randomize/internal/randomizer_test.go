package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomizeSimpleValues(t *testing.T) {

	t.Run("randomize string", func(t *testing.T) {
		type StringStruct struct {
			String string
		}
		got := randomize(StringStruct{})
	})

	t.Run("randomize integer", func(t *testing.T) {

	})

	t.Run("randomize float", func(t *testing.T) {

	})

	t.Run("randomize boolean", func(t *testing.T) {

	})

}
