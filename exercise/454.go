package main

import "fmt"

//Given four lists A, B, C, D of integer values, compute how many tuples (i, j, k, l) there are such that A[i] + B[j] + C[k] + D[l] is zero.
//
//To make problem a bit easier, all A, B, C, D have same length of N where 0 ≤ N ≤ 500. All integers are in the range of -228 to 228 - 1 and the result is guaranteed to be at most 231 - 1.
//
//Example:
//
//Input:
//A = [ 1, 2]
//B = [-2,-1]
//C = [-1, 2]
//D = [ 0, 2]
//
//Output:
//2
//
//Explanation:
//The two tuples are:
//1. (0, 0, 0, 1) -> A[0] + B[0] + C[0] + D[1] = 1 + (-2) + (-1) + 2 = 0
//2. (1, 1, 0, 0) -> A[1] + B[1] + C[0] + D[0] = 2 + (-1) + (-1) + 0 = 0

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var record = make(map[int]int)
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			record[C[i]+D[j]]++
		}
	}

	var res int
	for _, v := range A {
		for _, vv := range B {
			if cnt, ok := record[0-v-vv]; ok {
				res += cnt
			}
		}
	}

	return res
}

func main() {
	fmt.Println("expect 2, result:", fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}))
}
