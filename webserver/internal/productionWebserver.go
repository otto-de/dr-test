package internal

import (
	"drtest/generated"
)

type ProductionWebserver struct{}

func (server ProductionWebserver) GetRecordNames() []string {
	return generated.GetRecordNames()
}

func (server ProductionWebserver) GenerateEntity(recordName string, amount int) ([]interface{}, error) {
	return generated.Generate(recordName, amount)
}
