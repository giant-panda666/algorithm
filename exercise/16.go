package main

import (
	"fmt"
	"math"
	"sort"
)

//Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one solution.
//
//    For example, given array S = {-1 2 1 -4}, and target = 1.
//
//    The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	sum := nums[0] + nums[1] + nums[2]

	for i := 0; i < len(nums); i++ {
		for l, r := i+1, len(nums)-1; l < r; {
			tmpSum := nums[i] + nums[l] + nums[r]
			if math.Abs(float64(sum-target)) > math.Abs(float64(tmpSum-target)) {
				sum = tmpSum
			}
			if tmpSum == target {
				return target
			} else if tmpSum < target {
				l++
			} else {
				r--
			}
		}
	}

	return sum
}

func main() {
	var nums = []int{-1, 2, 1, -4}
	fmt.Println("expect 2, result:", threeSumClosest(nums, 1))
}
