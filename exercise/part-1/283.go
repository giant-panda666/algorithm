package main

import "fmt"

//Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].
//
//Note:
//    1.You must do this in-place without making a copy of the array.
//    2.Minimize the total number of operations.

func moveZeroes(nums []int) {
	i := -1 // nums[0, ... i] != 0
	j := 0  // nums[i+1, j-1] = 0
	for j < len(nums) {
		if nums[j] != 0 {
			i++
			if j != i {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		j++
	}
}

func main() {
	var nums = []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
