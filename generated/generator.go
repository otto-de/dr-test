package generated

import (
	"drtest/generated/person"
	"errors"
)

func Generate(recordName string, amount int) ([]interface{}, error) {
	switch recordName {
	case "person":
		return person.GeneratePerson(amount), nil
	default:
		return nil, errors.New("record not found")
	}
}
func GetRecordNames() []string {
	return []string{"person"}
}
