package main

import "fmt"

//Given a string, find the length of the longest substring without repeating characters.
//
//Examples:
//
//Given "abcabcbb", the answer is "abc", which the length is 3.
//
//Given "bbbbb", the answer is "b", with the length of 1.
//
//Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

func lengthOfLongestSubstring(s string) int {
	str := []rune(s)
	l, r := 0, 0 // str[l, ... r-1] 滑动窗口
	res := 0
	for r < len(str) {
		isDup, index := isDuplicateChar(str, l, r)
		if !isDup {
			if r-l+1 > res {
				res = r - l + 1
			}
			r++
		} else {
			l = index + 1
		}
	}

	return res
}

// 新加入的str[r]是否和[l, r-1]之间的字符重复，如果重复，则返回重复字符的位置
func isDuplicateChar(str []rune, l, r int) (bool, int) {
	for i := l; i < r; i++ {
		if str[i] == str[r] {
			return true, i
		}
	}
	return false, -1
}

func main() {
	var str1 = "abcabcbb"
	var str2 = "bbb"
	var str3 = "pwwkew"
	fmt.Println("expect 3, result", lengthOfLongestSubstring(str1))
	fmt.Println("expect 1, result", lengthOfLongestSubstring(str2))
	fmt.Println("expect 3, result", lengthOfLongestSubstring(str3))
}
