package main

import "fmt"

//Given two strings s and t, determine if they are isomorphic.
//
//Two strings are isomorphic if the characters in s can be replaced to get t.
//
//All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
//
//For example,
//Given "egg", "add", return true.
//
//Given "foo", "bar", return false.
//
//Given "paper", "title", return true.
//
//Note:
//You may assume both s and t have the same length.

func isIsomorphic(s string, t string) bool {
	rs, rt := []rune(s), []rune(t)
	if len(rs) != len(rt) {
		return false
	}

	var record1 = make(map[rune]rune)
	var record2 = make(map[rune]rune)
	for i := 0; i < len(rs); i++ {
		record1[rs[i]] = rt[i]
		record2[rt[i]] = rs[i]
	}
	for i := 0; i < len(rs); i++ {
		if rt[i] != record1[rs[i]] || rs[i] != record2[rt[i]] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("expect true, result:", isIsomorphic("egg", "add"))
	fmt.Println("expect false, reuslt:", isIsomorphic("foo", "bar"))
	fmt.Println("expect false, reuslt:", isIsomorphic("ab", "aa"))
	fmt.Println("expect true, result:", isIsomorphic("paper", "title"))
}
