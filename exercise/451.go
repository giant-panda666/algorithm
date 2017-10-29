package main

import (
	"fmt"
	"sort"
	"strings"
)

//Given a string, sort it in decreasing order based on the frequency of characters.
//
//Example 1:
//
//Input:
//"tree"
//
//Output:
//"eert"
//
//Explanation:
//'e' appears twice while 'r' and 't' both appear once.
//So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
//Example 2:
//
//Input:
//"cccaaa"
//
//Output:
//"cccaaa"
//
//Explanation:
//Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
//Note that "cacaca" is incorrect, as the same characters must be together.
//Example 3:
//
//Input:
//"Aabb"
//
//Output:
//"bbAa"
//
//Explanation:
//"bbaA" is also a valid answer, but "Aabb" is incorrect.
//Note that 'A' and 'a' are treated as two different characters.

func frequencySort(s string) string {
	var record = make(map[rune]int)
	for _, v := range s {
		record[v]++
	}
	var pairs []pair
	for k, v := range record {
		pairs = append(pairs, pair{k: string(k), v: v})
	}
	sort.Sort(PairSlice(pairs))

	var res string
	for _, v := range pairs {
		res += strings.Repeat(v.k, v.v)
	}
	return res
}

type pair struct {
	k string
	v int
}

type PairSlice []pair

func (p PairSlice) Len() int {
	return len(p)
}

func (p PairSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p PairSlice) Less(i, j int) bool {
	if p[i].v > p[j].v {
		return true
	}
	return false
}

func main() {
	fmt.Println("input tree, out:", frequencySort("tree"))
	fmt.Println("input cccaaa, out:", frequencySort("cccaaa"))
	fmt.Println("input Aabb, out:", frequencySort("Aabb"))
}
