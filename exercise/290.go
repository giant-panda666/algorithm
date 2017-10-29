package main

import (
	"fmt"
	"strings"
)

//Given a pattern and a string str, find if str follows the same pattern.
//
//Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.
//
//Examples:
//pattern = "abba", str = "dog cat cat dog" should return true.
//pattern = "abba", str = "dog cat cat fish" should return false.
//pattern = "aaaa", str = "dog cat cat dog" should return false.
//pattern = "abba", str = "dog dog dog dog" should return false.
//Notes:
//You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.

func wordPattern(pattern string, str string) bool {
	strs := strings.Fields(str)
	if len(pattern) != len(strs) {
		return false
	}

	var p2str = make(map[byte]string)
	var str2p = make(map[string]byte)
	for i := 0; i < len(pattern); i++ {
		str1, ok1 := p2str[pattern[i]]
		p2, ok2 := str2p[strs[i]]
		if ok1 != ok2 {
			return false
		}
		if ok1 {
			if str1 != strs[i] || p2 != pattern[i] {
				return false
			}
		} else {
			p2str[pattern[i]] = strs[i]
			str2p[strs[i]] = pattern[i]
		}
	}

	return true
}

func main() {
	fmt.Println("expect true, result:", wordPattern("abba", "dog cat cat dog"))
	fmt.Println("expect false, result:", wordPattern("abba", "dog cat cat fish"))
	fmt.Println("expect false, result:", wordPattern("aaaa", "dog cat cat dog"))
	fmt.Println("expect false, result:", wordPattern("abba", "dog dog dog dog"))
	fmt.Println("expect false, result:", wordPattern("aba", "dog cat cat"))
}
