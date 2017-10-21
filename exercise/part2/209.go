package main

import "fmt"

//Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.
//
//For example, given the array [2,3,1,2,4,3] and s = 7,
//the subarray [4,3] has the minimal length under the problem constraint.
//
//click to show more practice.
//
//Credits:
//Special thanks to @Freezen for adding this problem and creating all test cases.

func minSubArrayLen(s int, nums []int) int {
	l, r := 0, -1 // nums[l, ... r]滑动窗口
	sum, res := 0, len(nums)+1
	for l < len(nums) {
		if sum < s && r+1 < len(nums) {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= s && res > r-l+1 {
			res = r - l + 1
		}
	}

	if res == len(nums)+1 {
		res = 0
	}

	return res
}

func main() {
	var nums = []int{2, 3, 1, 2, 4, 3}
	fmt.Println("expect 2, result", minSubArrayLen(7, nums))
}
