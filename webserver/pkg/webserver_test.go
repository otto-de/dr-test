package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testWebserver struct {

}
func (t testWebserver) GetRecordNames() []string {
	return []string{""}
}

func (t testWebserver) GenerateEntity(recordName string, amount int) (interface{}, error) {
	return defaultEntity, nil
}

type testEntity struct {
	Name string
}
var defaultEntity = testEntity{"foo"}

func TestStart(t *testing.T) {
	recordName := "foo"
	t.Run("response contains entity", func (t *testing.T) {
		// given
		// recordName

		// when
		entity, _ := getResponse(testWebserver{}, recordName, 1)

		// then
		assert.Equal(t, entity.Body, defaultEntity)
	})

	t.Run("Headers contain Content-Type", func(t *testing.T) {
		// given
		// recordName

		// when
		entity, _ := getResponse(testWebserver{}, recordName, 1)

		// then
		assert.Equal(t, entity.Headers["Content-Type"], "application/json")
	})
}
