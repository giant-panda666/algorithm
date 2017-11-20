package main

import "fmt"

//Given a set of distinct integers, nums, return all possible subsets (the power set).
//
//Note: The solution set must not contain duplicate subsets.
//
//For example,
//If nums = [1,2,3], a solution is:
//
//[
//  [3],
//  [1],
//  [2],
//  [1,2,3],
//  [1,3],
//  [2,3],
//  [1,2],
//  []
//]
//

func subsets(nums []int) [][]int {
	var cur []int
	var res [][]int
	generateSubset(nums, 0, &cur, &res)
	return res
}

func generateSubset(nums []int, index int, cur *[]int, res *[][]int) {
	tmp := make([]int, len(*cur))
	copy(tmp, *cur)
	*res = append(*res, tmp)

	for i := index; i < len(nums); i++ {
		*cur = append(*cur, nums[i])
		generateSubset(nums, i+1, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}

func main() {
	var nums = []int{1, 2, 3, 4}
	fmt.Println(subsets(nums))
}
