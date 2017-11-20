package main

import "fmt"

//Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).
//
//Note: The solution set must not contain duplicate subsets.
//
//For example,
//If nums = [1,2,2], a solution is:
//
//[
//  [2],
//  [1],
//  [1,2,2],
//  [2,2],
//  [1,2],
//  []
//]

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var cur []int
	var res [][]int
	generateSubsetWithDup(nums, 0, &cur, &res)
	return res

}

func generateSubsetWithDup(nums []int, index int, cur *[]int, res *[][]int) {
	tmp := make([]int, len(*cur))
	copy(tmp, *cur)
	*res = append(*res, tmp)

	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		*cur = append(*cur, nums[i])
		generateSubsetWithDup(nums, i+1, cur, res)
		*cur = (*cur)[:len(*cur)-1]
	}
}

func main() {
	var nums = []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
