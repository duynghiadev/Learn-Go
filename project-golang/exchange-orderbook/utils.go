package main

import (
	"math/rand"
	"time"
)

func randUser() int64 {
	return int64(randNumber(10000000))
}

func randNumber(n int) int {
	source := rand.NewSource(time.Now().UnixNano())
	num := rand.New(source)

	return num.Intn(n)
}
