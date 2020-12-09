package generated

import (
	"errors"
)

func Generate(recordName string, amount int) ([]interface{}, error) {
	switch recordName {
	default:
		return nil, errors.New("record not found")
	}
}
func GetRecordNames() []string {
	return []string{"person"}
}
