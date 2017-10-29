package main

import (
	"fmt"
)

//Write an algorithm to determine if a number is "happy".
//
//A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.
//
//Example: 19 is a happy number
//
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1

func isHappy(n int) bool {
	var record = make(map[int]bool)
	for !record[n] {
		record[n] = true
		sum := 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		if sum == 1 {
			return true
		}
		n = sum
	}
	return false
}

func main() {
	fmt.Println("expect true, result:", isHappy(19))
}
