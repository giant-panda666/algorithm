package main

import "fmt"

//Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
//
//For example,
//If n = 4 and k = 2, a solution is:
//
//[
//  [2,4],
//  [3,4],
//  [2,3],
//  [1,2],
//  [1,3],
//  [1,4],
//]

func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return nil
	}
	var cur []int
	var res [][]int
	generateCombine(n, k, 1, &cur, &res)
	return res
}

func generateCombine(n, k, index int, cur *[]int, res *[][]int) {
	if len(*cur) == k {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for j := index; j <= n-(k-len(*cur))+1; j++ {
		*cur = append(*cur, j)
		generateCombine(n, k, j+1, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
