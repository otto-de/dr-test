package api

import "drtest/randomize/internal"

func randomize(strukt interface{}, configuration internal.Configuration) interface{} {
	return internal.Randomize(strukt, configuration)
}
