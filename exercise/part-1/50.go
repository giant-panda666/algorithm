package main

import "fmt"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	} else {
		return x * myPow(x*x, n/2)
	}
}

func main() {
	fmt.Println("-2.0, 3", myPow(-2.0, 3))
	fmt.Println("2.0, 3", myPow(2.0, 3))
	fmt.Println("2.0, -3", myPow(2.0, -3))
	fmt.Println("2.0, 0", myPow(2.0, 0))
}
