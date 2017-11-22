package main

import "fmt"

//Given a positive integer n, break it into the sum of at least two positive integers and maximize the product of those integers. Return the maximum product you can get.
//
//For example, given n = 2, return 1 (2 = 1 + 1); given n = 10, return 36 (10 = 3 + 3 + 4).
//
//Note: You may assume that n is not less than 2 and not larger than 58.

func integerBreak(n int) int {
	// n is not less than 2
	var product = make([]int, n+1)
	product[2] = 1
	for i := 3; i < len(product); i++ {
		var max int
		for j := 1; j <= i/2; j++ {
			var a, b = j, i - j
			if j > 2 && product[j] > j {
				a = product[j]
			}
			if i-j > 2 && product[i-j] > i-j {
				b = product[i-j]
			}
			if a*b > max {
				max = a * b
			}
		}
		product[i] = max
	}
	return product[n]
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}
