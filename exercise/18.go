package main

import (
	"fmt"
	"sort"
)

//Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
//Note: The solution set must not contain duplicate quadruplets.
//
//For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A solution set is:
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}

	var res [][]int
	sort.Ints(nums)

	if 4*nums[0] > target || 4*nums[len(nums)-1] < target {
		return nil
	}

	for a := 0; a < len(nums); a++ {
		tmp := threeSum(nums[a+1:], target-nums[a])
		for _, v := range tmp {
			res = append(res, []int{nums[a], v[0], v[1], v[2]})
		}
		for a+1 < len(nums) && nums[a] == nums[a+1] {
			a++
		}
	}

	return res
}

func threeSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	if 3*nums[0] > target || 3*nums[len(nums)-1] < target {
		return nil
	}

	var res [][]int
	for a := 0; a < len(nums); a++ {
		tmp := twoSum(nums[a+1:], target-nums[a])
		for _, v := range tmp {
			res = append(res, []int{nums[a], v[0], v[1]})
		}
		for a+1 < len(nums) && nums[a] == nums[a+1] {
			a++
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	if len(nums) < 2 {
		return nil
	}

	if 2*nums[0] > target || 2*nums[len(nums)-1] < target {
		return nil
	}

	var res [][]int
	l, r := 0, len(nums)-1
	for l < r {
		sum := nums[l] + nums[r]
		if sum == target {
			res = append(res, []int{nums[l], nums[r]})
			for l < r && nums[l] == nums[l+1] {
				l++
			}

			for l < r && nums[r] == nums[r-1] {
				r--
			}
			l++
			r--
		} else if sum < target {
			l++
		} else if sum > target {
			r--
		}
	}
	return res
}

func main() {
	var nums = []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}
