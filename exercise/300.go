package main

//Given an unsorted array of integers, find the length of longest increasing subsequence.
//
//For example,
//Given [10, 9, 2, 5, 3, 7, 101, 18],
//The longest increasing subsequence is [2, 3, 7, 101], therefore the length is 4. Note that there may be more than one LIS combination, it is only necessary for you to return the length.
//
//Your algorithm should run in O(n2) complexity.
//
//Follow up: Could you improve it to O(n log n) time complexity?

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = 1
	}

	var max = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && res[j]+1 > res[i] {
				res[i] = res[j] + 1
			}
		}
		if res[i] > max {
			max = res[i]
		}
	}

	return max
}
