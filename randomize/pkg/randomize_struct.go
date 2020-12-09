package api

import (
	"drtest/randomize/api"
	"drtest/randomize/internal"
)

func RandomizeWithDefaults(strukt interface{}) interface{} {
	config := api.Configuration{
		MaxListLength:   10,
		MaxStringLength: 10,
	}
	return internal.Randomize(strukt, config)
}

func Randomize(strukt interface{}, configuration api.Configuration) interface{} {
	return internal.Randomize(strukt, configuration)
}
