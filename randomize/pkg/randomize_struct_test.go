package api

import (
	"drtest/randomize/api"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFillPrimitiveAvroFields(t *testing.T) {
	type Avro struct {
		Bool   bool
		Int    int32
		Long   int64
		Float  float32
		Double float64
		Bytes  []byte
		String string
		Array  []string
		Map    map[string]int
	}

	avro := Randomize(&Avro{}, api.Configuration{
		MinListLength:   1,
		MaxListLength:   20,
		MaxStringLength: 2,
	}).(*Avro)
	assert.NotNil(t, &avro)
	assert.NotEmpty(t, avro.Bytes, "bytes array is empty")
	assert.NotEmpty(t, avro.Array, "string array is empty")
	assert.NotNil(t, avro.Map, "map is empty")

}

func TestRandomize(t *testing.T) {

	randomized := Randomize(&Person{}, api.Configuration{
		MaxListLength:   4,
		MaxStringLength: 5,
	})
	assert.NotNil(t, randomized.(*Person))
}

type Pet struct {
	Name string
	Age  int64
}

type Coordinates struct {
	Lat float64
	Lon float64
}

type Person struct {
	FirstName         string
	LastName          string
	Hobbies           []string
	LuckyNumbers      []int64
	Cool              bool
	Balance           float64
	Coordinates       Coordinates
	Pets              []Pet
	SuperBesteFreunde []Person
}

func (p Person) String() string {
	indent, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err.Error()
	}
	return string(indent)
}
