package api

import (
	"drtest/randomize/api"
	"drtest/randomize/internal"
)

func RandomizeWithDefaults(strukt interface{}) interface{} {
	return internal.Randomize(strukt, api.DefaultConfiguration())
}

func Randomize(strukt interface{}, configuration api.Configuration) interface{} {
	return internal.Randomize(strukt, configuration)
}
