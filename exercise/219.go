package main

import "fmt"

//Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such that nums[i] = nums[j] and the absolute difference between i and j is at most k.

func containsNearbyDuplicate(nums []int, k int) bool {
	var record = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true
		}
		record[nums[i]]++
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}

func main() {
	var nums = []int{1}
	fmt.Println("expect false, result:", containsNearbyDuplicate(nums, 1))
}
