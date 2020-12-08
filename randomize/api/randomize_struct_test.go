package api

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomize(t *testing.T) {

	randomized := randomize(&Person{})
	person := randomized.(*Person)
	fmt.Println(person)
	assert.True(t, len(person.FirstName) > 0)
	assert.True(t, len(person.LastName) > 0)
	assert.True(t, len(person.Hobbies) > 0)
	assert.True(t, len(person.Hobbies[0]) > 0)
	assert.True(t, person.Balance != 0)
}

type Person struct {
	FirstName    string
	LastName     string
	Hobbies      []string
	LuckyNumbers []int64
	Cool         bool
	Balance      float64
}

func (p Person) String() string {
	indent, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(indent)
}
