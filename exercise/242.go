package main

import "fmt"

//Given two strings s and t, write a function to determine if t is an anagram of s.
//
//For example,
//s = "anagram", t = "nagaram", return true.
//s = "rat", t = "car", return false.
//
//Note:
//You may assume the string contains only lowercase alphabets.
//
//Follow up:
//What if the inputs contain unicode characters? How would you adapt your solution to such case?

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sl, tl := []rune(s), []rune(t)
	var recored = make(map[rune]int)
	for i := 0; i < len(sl); i++ {
		recored[sl[i]]++
		recored[tl[i]]--
	}
	for _, v := range recored {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("expect true, result:", isAnagram("anagram", "nagaram"))
	fmt.Println("expect false, result:", isAnagram("cat", "car"))
}
