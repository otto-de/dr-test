package internal

import (
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString() string {
	b := make([]byte, 1+rand.Int31n(100))
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func randomInt() int64 {
	return rand.Int63()
}

func randomFloat() float64 {
	return rand.Float64()
}

func randomBool() bool {
	return rand.Float32() <= 0.5
}
