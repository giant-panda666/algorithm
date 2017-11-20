package main

import "fmt"

//Find all possible combinations of k numbers that add up to a number n, given that only numbers from 1 to 9 can be used and each combination should be a unique set of numbers.
//
//
//Example 1:
//
//Input: k = 3, n = 7
//
//Output:
//
//[[1,2,4]]
//
//Example 2:
//
//Input: k = 3, n = 9
//
//Output:
//
//[[1,2,6], [1,3,5], [2,3,4]]

func combinationSum3(k int, n int) [][]int {
	if k <= 0 || n <= 0 || k > n {
		return nil
	}
	var cur []int
	var curSum int
	var res [][]int
	generateCombinationSum3(k, n, 1, curSum, &cur, &res)
	return res
}

func generateCombinationSum3(k, n int, index int, curSum int, cur *[]int, res *[][]int) {
	if n == curSum && len(*cur) == k {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := index; i <= 9-(k-len(*cur))+1; i++ {
		if curSum+index > n {
			return
		}
		curSum += i
		*cur = append(*cur, i)
		generateCombinationSum3(k, n, i+1, curSum, cur, res)
		curSum -= i
		*cur = (*cur)[:len(*cur)-1]
	}
}

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
}
