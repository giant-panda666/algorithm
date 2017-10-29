package main

import "fmt"

//Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
//
//Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
//
//The order of output does not matter.
//
//Example 1:
//
//Input:
//s: "cbaebabacd" p: "abc"
//
//Output:
//[0, 6]
//
//Explanation:
//The substring with start index = 0 is "cba", which is an anagram of "abc".
//The substring with start index = 6 is "bac", which is an anagram of "abc".
//Example 2:
//
//Input:
//s: "abab" p: "ab"
//
//Output:
//[0, 1, 2]
//
//Explanation:
//The substring with start index = 0 is "ab", which is an anagram of "ab".
//The substring with start index = 1 is "ba", which is an anagram of "ab".
//The substring with start index = 2 is "ab", which is an anagram of "ab".

func findAnagrams(s string, p string) []int {
	sr, pr := []rune(s), []rune(p)
	if len(sr) < len(pr) {
		return nil
	}

	var res []int
	var count int
	var pl = make(map[rune]int)
	for i := 0; i < len(pr); i++ {
		pl[pr[i]]++
		count++
	}

	// cbaebabacd abc
	for l, r := 0, 0; r < len(sr); {
		if cn, ok := pl[sr[r]]; !ok {
			for l < r {
				pl[sr[l]]++
				count++
				l++
			}
			l, r = r+1, r+1
		} else {
			if cn > 0 {
				pl[sr[r]]--
				count--
				if count == 0 {
					res = append(res, r-len(pr)+1)
					pl[sr[l]]++
					count++
					l++
				}
				r++
			} else {
				pl[sr[l]]++
				count++
				l++
			}
		}
	}

	return res
}

func main() {
	fmt.Println("expect [0, 6], result:", findAnagrams("cbaebabacd", "abc"))
}
