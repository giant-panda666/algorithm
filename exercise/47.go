package main

import "fmt"

//Given a collection of numbers that might contain duplicates, return all possible unique permutations.
//
//For example,
//[1,1,2] have the following unique permutations:
//[
//  [1,1,2],
//  [1,2,1],
//  [2,1,1]
//]
//

func permuteUnique(nums []int) [][]int {
	var cur []int
	var res [][]int
	var used = make([]bool, len(nums))
	generatePermuteUnique(nums, 0, used, &cur, &res)
	return res
}

func generatePermuteUnique(nums []int, index int, used []bool, cur *[]int, res *[][]int) {
	if len(nums) == index {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	var selected = make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if !used[i] && !selected[nums[i]] {
			*cur = append(*cur, nums[i])
			used[i] = true
			selected[nums[i]] = true
			generatePermuteUnique(nums, index+1, used, cur, res)
			*cur = (*cur)[:len(*cur)-1]
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
