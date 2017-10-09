package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	var mh = newMaxHeap(10)
	for i := 0; i < 10; i++ {
		mh.Insert(rand.Int() % 100)
	}

	for !mh.isEmpty() {
		fmt.Print(mh.ExtractMax(), " ")
	}
	fmt.Println()
}
