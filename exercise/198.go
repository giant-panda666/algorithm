package main

//You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.
//
//Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var res = make([]int, len(nums)+1)
	res[0] = 0
	res[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		if res[i-1] > res[i-2]+nums[i-1] {
			res[i] = res[i-1]
		} else {
			res[i] = res[i-2] + nums[i-1]
		}
	}
	return res[len(nums)]
}
