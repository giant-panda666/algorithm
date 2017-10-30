package main

import "fmt"

//Given an array of strings, group anagrams together.
//
//For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
//Return:
//
//[
//  ["ate", "eat","tea"],
//  ["nat","tan"],
//  ["bat"]
//]
//Note: All inputs will be in lower-case.

func groupAnagrams(strs []string) [][]string {
	prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}
	var record = make(map[int][]string)
	for _, v := range strs {
		var product = 1
		for _, vv := range v {
			product *= prime[vv-'a']
		}
		record[product] = append(record[product], v)
	}

	var res [][]string
	for _, v := range record {
		res = append(res, v)
	}

	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
