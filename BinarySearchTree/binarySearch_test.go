package main

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var a1 = []int{1, 2, 3, 4, 5, 6, 6, 7, 7, 8}
	var a2 = []int{1, 2, 3, 4, 5, 7, 8, 9}
	var a3 = []int{1, 2, 3, 4, 5, 7, 8, 9, 10}
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a1 6 binarySearch", binarySearch(a1, 6))
	fmt.Println("a2 6 binarySearch", binarySearch(a2, 6))
	fmt.Println("a1 6 floor", floor(a1, 6))
	fmt.Println("a1 6 ceil", ceil(a1, 6))
	fmt.Println("a2 6 floor", floor(a2, 6))
	fmt.Println("a2 6 ceil", ceil(a2, 6))
	fmt.Println("a3 floor", floor(a3, 6))
	fmt.Println("a3 6 ceil", ceil(a3, 6))
}
