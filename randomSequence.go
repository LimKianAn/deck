package main

import (
	"math/rand"
	"time"
)

func newRandomSequence(size int) []int {
	x := make([]int, size)
	for i := range x {
		x[i] = i + 1
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(x), func(i, j int) { x[i], x[j] = x[j], x[i] })

	return x
}
