package api

import "drtest/randomize/internal"

func Randomize(strukt interface{}, configuration internal.Configuration) interface{} {
	return internal.Randomize(strukt, configuration)
}
