package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isAscii(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func isPalindrome(s string) bool {
	tmp := []rune(strings.ToUpper(s))
	for i, j := 0, len(tmp)-1; j > i; {
		for i < j && !unicode.IsDigit(tmp[i]) && !isAscii(tmp[i]) {
			i++
		}
		for i < j && !unicode.IsDigit(tmp[j]) && !isAscii(tmp[j]) {
			j--
		}
		if tmp[i] != tmp[j] {
			return false
		}
		i, j = i+1, j-1
	}
	return true
}

func main() {
	var s1 = "A man, a plan, a canal: Panama"
	var s2 = "race a car"
	var s3 = "a."
	var s4 = "0P"
	fmt.Println(isPalindrome(s1))
	fmt.Println(isPalindrome(s2))
	fmt.Println(isPalindrome(s3))
	fmt.Println(isPalindrome(s4))
}
