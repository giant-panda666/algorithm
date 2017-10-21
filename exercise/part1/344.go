package main

import "fmt"

//Write a function that takes a string as input and returns the string reversed.
//
//Example:
//Given s = "hello", return "olleh".

func reverseString(s string) string {
	tmp := []rune(s)
	for i, j := 0, len(tmp)-1; j > i; i, j = i+1, j-1 {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	}
	return string(tmp)
}

func main() {
	var s = "hello"
	fmt.Println(reverseString(s))
}
