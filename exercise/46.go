package main

import "fmt"

//Given a collection of distinct numbers, return all possible permutations.
//
//For example,
//[1,2,3] have the following permutations:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
//
func permute(nums []int) [][]int {
	var cur []int
	var res [][]int
	var used = make([]bool, len(nums))
	generatePermute(nums, 0, used, &cur, &res)
	return res
}

func generatePermute(nums []int, index int, used []bool, cur *[]int, res *[][]int) {
	if len(nums) == index {
		tmp := make([]int, len(*cur))
		copy(tmp, *cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			*cur = append(*cur, nums[i])
			used[i] = true
			generatePermute(nums, index+1, used, cur, res)
			*cur = (*cur)[:len(*cur)-1]
			used[i] = false
		}
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
