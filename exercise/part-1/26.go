package main

import "fmt"

//Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
//
//Do not allocate extra space for another array, you must do this in place with constant memory.
//
//For example,
//Given input array nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	// nums[0, ... i] none duplicate elements
	// nums[i+1, ... j] duplicate elements
	i, j, k := 0, 0, 1
	for k < len(nums) {
		if nums[k] != nums[k-1] {
			i++
			nums[i] = nums[k]
		}
		j++
		k++
	}
	return i + 1
}

func main() {
	var nums1 = []int{1, 1, 2}
	var nums2 = []int{1, 2, 2, 3, 4, 4}
	fmt.Println(removeDuplicates(nums1), nums1)
	fmt.Println(removeDuplicates(nums2), nums2)
}
