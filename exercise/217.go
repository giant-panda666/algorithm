package main

import "fmt"

//Given an array of integers, find if the array contains any duplicates. Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

func containsDuplicate(nums []int) bool {
	var record = make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true
		}
		record[nums[i]] = struct{}{}
	}
	return false
}

func main() {
	var nums = []int{1, 1}
	fmt.Println("expect true, result:", containsDuplicate(nums))
}
