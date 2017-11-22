package main

import (
	"fmt"
	"strconv"
)

//A message containing letters from A-Z is being encoded to numbers using the following mapping:
//
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//Given an encoded message containing digits, determine the total number of ways to decode it.
//
//For example,
//Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).
//
//The number of ways decoding "12" is 2.

func numDecodings(s string) int {
	if s == "" {
		return 0
	}

	var res = make([]int, len(s)+1)
	res[0] = 1
	if s[0] == '0' {
		res[1] = 0
	} else {
		res[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		first, _ := strconv.Atoi(s[i-1 : i])
		second, _ := strconv.Atoi(s[i-2 : i])
		if first >= 1 && first <= 9 {
			res[i] += res[i-1]
		}
		if second >= 10 && second <= 26 {
			res[i] += res[i-2]
		}
	}

	return res[len(s)]
}

func main() {
	fmt.Println(numDecodings("12"))
}
