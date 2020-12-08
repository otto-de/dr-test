package api

import "drtest/randomize/internal"

func randomize(strukt interface{}) interface{} {
	return internal.Randomize(strukt, internal.Configuration{
		MaxListSize:     10,
		MaxStringLength: 5,
	})
}
