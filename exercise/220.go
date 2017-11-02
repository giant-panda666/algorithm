package main

import (
	"fmt"
)

//Given an array of integers, find out whether there are two distinct indices i and j in the array such that the absolute difference between nums[i] and nums[j] is at most t and the absolute difference between i and j is at most k.

// |nums[i]-nums[j]| <= t => |nums[i]-nums[j]| < t+1
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if len(nums) <= 1 || k < 0 || t < 0 {
		return false
	}
	m := make(map[int]int)
	w := t + 1
	for i := 0; i < len(nums); i++ {
		s := nums[i] / w
		if nums[i] < 0 {
			s = s - 1
		}

		if _, ok := m[s]; ok {
			return true
		}

		if num, ok := m[s-1]; ok && abs(nums[i], num) < w {
			return true
		}

		if num, ok := m[s+1]; ok && abs(nums[i], num) < w {
			return true
		}
		m[s] = nums[i]
		if i >= k {
			delete(m, nums[i-k]/w)
		}
	}

	return false
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	var nums = []int{-3, 3}
	k, t := 2, 4
	fmt.Println("expect false, result:", containsNearbyAlmostDuplicate(nums, k, t))
}
