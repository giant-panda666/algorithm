package main

import "fmt"

//Given a string s, partition s such that every substring of the partition is a palindrome.
//
//Return all possible palindrome partitioning of s.
//
//For example, given s = "aab",
//Return
//
//[
//  ["aa","b"],
//  ["a","a","b"]
//]
func partition(s string) [][]string {
	var cur []string
	var res [][]string
	partitionImpl(s, 0, &cur, &res)
	return res
}

func partitionImpl(s string, l int, cur *[]string, res *[][]string) {
	if len(s) == l && *cur != nil {
		var tmp = make([]string, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		fmt.Println("l:", l, "*cur:", *cur)
	}
	for i := l; i < len(s); i++ {
		if isPalindrome(s[l : i+1]) {
			*cur = append(*cur, s[l:i+1])
			partitionImpl(s, i+1, cur, res)
			*cur = (*cur)[:len(*cur)-1]
		}
	}
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("aaa"))
}
