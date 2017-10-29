package main

import "fmt"

//Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).
//
//For example,
//S = "ADOBECODEBANC"
//T = "ABC"
//Minimum window is "BANC".
//
//Note:
//If there is no such window in S that covers all characters in T, return the empty string "".
//
//If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.

func minWindow(s string, t string) string {
	sr, tr := []rune(s), []rune(t)
	if len(s) < len(t) {
		return ""
	}

	var resl, resr = 0, len(sr) + 1
	var l, r = 0, 0
	var count = 0
	var tm = make(map[rune]int)
	for i := 0; i < len(tr); i++ {
		if _, ok := tm[tr[i]]; !ok {
			count++
		}
		tm[tr[i]]++
	}

	for r < len(sr) {
		if cn, ok := tm[sr[r]]; ok {
			tm[sr[r]]--
			if cn-1 == 0 {
				count--
			}
		}
		r++

		if count == 0 {
			for l < r {
				if cn, ok := tm[sr[l]]; ok {
					if cn < 0 {
						tm[sr[l]]++
						l++
					}
					if cn == 0 {
						if resr-resl > r-l {
							resl, resr = l, r
							fmt.Println(string(sr[resl:resr]))
						}
						tm[sr[l]]++
						l++
						count++
						for l < r {
							if _, ok := tm[sr[l]]; ok {
								break
							}
							l++
						}
						break
					}
				} else {
					l++
				}
			}
		}
	}
	if resr-resl != len(sr)+1 {
		return string(sr[resl:resr])
	}
	return ""
}

func main() {
	fmt.Println("expect BANC, result:", minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println("expect aa, result:", minWindow("aa", "aa"))
}
