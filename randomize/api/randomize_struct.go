package api

import "drtest/randomize/internal"

func RandomizeWithDefaults(strukt interface{}) interface{} {
	config := Configuration{
		MaxListSize:     10,
		MaxStringLength: 10,
	}
	return internal.Randomize(strukt, config)
}

func Randomize(strukt interface{}, configuration Configuration) interface{} {
	return internal.Randomize(strukt, configuration)
}

type Configuration struct {
	MaxListSize     int
	MaxStringLength int
}
