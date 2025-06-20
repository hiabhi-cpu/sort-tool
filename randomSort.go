package main

import (
	"math/rand"
	"time"
)

func randomSort(slice []string) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Shuffle the slice
	r.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}
