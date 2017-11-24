package main

import "fmt"

//Given a non-empty array containing only positive integers, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.
//
//Note:
//Each of the array element will not exceed 100.
//The array size will not exceed 200.
//Example 1:
//
//Input: [1, 5, 11, 5]
//
//Output: true
//
//Explanation: The array can be partitioned as [1, 5, 5] and [11].
//Example 2:
//
//Input: [1, 2, 3, 5]
//
//Output: false
//
//Explanation: The array cannot be partitioned into equal sum subsets.

var memo [][]int

func canPartition(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 || sum < 0 {
		return false
	}

	memo = make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		memo[i] = make([]int, sum/2+1)
	}

	return tryPartition(nums, len(nums)-1, sum/2)
}

func tryPartition(nums []int, index, sum int) bool {
	if sum < 0 || index < 0 {
		return false
	}
	if sum == 0 {
		memo[index][sum] = 1
		return true
	}

	if memo[index][sum] != 0 {
		return memo[index][sum] == 1
	}

	if tryPartition(nums, index-1, sum) || tryPartition(nums, index-1, sum-nums[index]) {
		memo[index][sum] = 1
	} else {
		memo[index][sum] = -1
	}
	return memo[index][sum] == 1
}

func main() {
	var nums = []int{1, 1}
	fmt.Println(canPartition(nums))
}
