package generated

import (
	"drtest/generated/person"
	"errors"
)

func Generate(structName string, amount int) ([]interface{}, error) {
	switch structName {
	case "person":
		return person.GeneratePerson(amount), nil
	default:
		return nil, errors.New("struct not found")
	}
}
