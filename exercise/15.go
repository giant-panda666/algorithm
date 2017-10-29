package main

import (
	"fmt"
	"sort"
)

//Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
//Note: The solution set must not contain duplicate triplets.
//
//For example, given array S = [-1, 0, 1, 2, -1, -4],
//
//A solution set is:
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	var res [][]int
	for i := 0; i < len(nums); i++ {
		for k, j := i+1, len(nums)-1; k < j; {
			if nums[i]+nums[k]+nums[j] < 0 {
				k++
			} else if nums[i]+nums[k]+nums[j] > 0 {
				j--
			} else {
				res = append(res, []int{nums[i], nums[k], nums[j]})
				for k < j && nums[k] == nums[k+1] {
					k++
				}
				for k < j && nums[j] == nums[j-1] {
					j--
				}
				k++
				j--
			}
		}

		for ; i < len(nums)-1 && nums[i] == nums[i+1]; i++ {
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{-2, 0, 0, 2, 2}))
}
