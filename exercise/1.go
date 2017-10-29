package main

import "fmt"

//Given an array of integers, return indices of the two numbers such that they add up to a specific target.
//
//You may assume that each input would have exactly one solution, and you may not use the same element twice.
//
//Example:
//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	var record = make(map[int]int)
	for i, v := range nums {
		if index, ok := record[target-nums[i]]; ok {
			return []int{index, i}
		}
		record[v] = i
	}

	return nil
}

func main() {
	var nums = []int{2, 7, 11, 15}
	var target = 9
	fmt.Println("expect [0, 1], reuslt:", twoSum(nums, target))
}
