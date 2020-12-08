package api

import "drtest/randomize/internal"

func randomize(strukt interface{}) interface{} {
	return internal.Randomize(strukt)
}
