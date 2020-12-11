package randomContent

import (
	"drtest/randomize/api"
	"math/rand"
	"reflect"
)

func EmptyMap(p reflect.Type, configuration api.Configuration) (reflect.Value, int) {
	mapType := reflect.MapOf(reflect.TypeOf("str"), p)
	size := getLen(configuration.MinMapLength, configuration.MaxMapLength)
	return reflect.MakeMapWithSize(mapType, size), size

}

func EmptySlice(p reflect.Type, configuration api.Configuration) (reflect.Value, int) {
	sliceOfType := reflect.SliceOf(p)
	size := getLen(configuration.MinListLength, configuration.MaxListLength)
	return reflect.MakeSlice(sliceOfType, 0, size), size
}

func getLen(minLen, maxLen int) int {
	if maxLen <= 0 {
		return 1
	}
	proposed := rand.Intn(maxLen)
	if proposed < minLen {
		return minLen
	}
	return proposed
}
