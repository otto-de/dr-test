package pkg

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testWebserver struct {
}

func (t testWebserver) GetRecordNames() []string {
	return []string{""}
}

func (t testWebserver) GenerateEntity(recordName string, amount int) ([]interface{}, error) {
	entities := make([]interface{}, amount)
	for i := 0; i < amount; i++ {
		entities[i] = defaultEntity
	}
	fmt.Printf("entities: %#v\n", entities)
	return entities, nil
}

type testEntity struct {
	Name string
}

var defaultEntity = testEntity{"foo"}

func TestStart(t *testing.T) {
	recordName := "foo"

	t.Run("response contains entity", func(t *testing.T) {
		// given
		// recordName

		// when
		entity, _ := getResponse(testWebserver{}, recordName, 1)

		// then
		assert.Equal(t, entity.Body, []interface{}{defaultEntity})
	})

	t.Run("Headers contain Content-Type", func(t *testing.T) {
		// given
		// recordName

		// when
		entity, _ := getResponse(testWebserver{}, recordName, 1)

		// then
		assert.Equal(t, entity.Headers["Content-Type"], "application/json")
	})

	t.Run("Multiple elements are created", func(t *testing.T) {
		// given
		// recordName

		// when
		entities, _ := getResponse(testWebserver{}, recordName, 5)

		data := entities.Body.([]interface{})
		assert.True(t, len(data) == 5)
	})
}
