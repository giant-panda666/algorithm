package main

import "fmt"

//Given an array and a value, remove all instances of that value in place and return the new length.
//
//Do not allocate extra space for another array, you must do this in place with constant memory.
//
//The order of elements can be changed. It doesn't matter what you leave beyond the new length.
//
//Example:
//Given input array nums = [3,2,2,3], val = 3
//
//Your function should return length = 2, with the first two elements of nums being 2.

func removeElement(nums []int, val int) int {
	length := len(nums)
	i, j := 0, len(nums)-1
	for {
		for i < len(nums) && nums[i] != val {
			i++
		}
		for j >= 0 && nums[j] == val {
			j--
		}

		if j >= i {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			nums = nums[:j+1]
			return length - j - 1
		}
	}
}

func main() {
	var nums = []int{3, 2, 2, 3}
	var val = 3
	fmt.Println("expect 2, result is", removeElement(nums, val))
}
