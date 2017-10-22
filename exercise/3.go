package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring1(s string) int {
	var freq [256]int
	l, r := 0, 0
	res := 0

	str := []rune(s)
	for r < len(str) {
		if freq[str[r]] == 0 {
			freq[str[r]]++
			r++
		} else {
			freq[str[l]]--
			l++
		}

		if res < r-l {
			res = r - l
		}
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	var freq [256]int
	l, r := 0, -1
	res := 0

	str := []rune(s)
	for r+1 < len(str) {
		for r+1 < len(str) && freq[str[r+1]] == 0 {
			r++
			freq[str[r]]++
		}

		if res < r-l+1 {
			res = r - l + 1
		}

		if r+1 < len(str) {
			r++
			freq[str[r]]++
			for l <= r && freq[str[r]] == 2 {
				freq[str[l]]--
				l++
			}
		}
	}
	return res
}

func isDuplicateChar(str []rune, l, r int) int {
	for i := l; i < r; i++ {
		if str[i] == str[r] {
			return i
		}
	}
	return -1
}

func lengthOfLongestSubstring3(s string) int {
	l, r := 0, 0
	res := 0
	str := []rune(s)

	for r < len(str) {
		index := isDuplicateChar(str, l, r)
		if index != -1 {
			l = index + 1
		}

		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}

	return res
}

func testPerformance(algoName string, f func(string) int, s string) {
	start := time.Now()
	res := f(s)
	t := time.Since(start)

	fmt.Printf("%s: res=%d, time=%v\n", algoName, res, t)
}

func main() {
	n := 10000000
	var s []rune
	for i := 0; i < n; i++ {
		//s = append(s, rune(rand.Int()%26+97))
		s = append(s, rune(rand.Int()%95+32))
	}
	testPerformance("algo1", lengthOfLongestSubstring1, string(s))
	testPerformance("algo2", lengthOfLongestSubstring2, string(s))
	testPerformance("algo3", lengthOfLongestSubstring3, string(s))
}
