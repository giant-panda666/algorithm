package main

import "fmt"

//Given an array with n objects colored red, white or blue, sort them so that objects of the same color are adjacent, with the colors in the order red, white and blue.
//
//Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

//func sortColors(nums []int) {
//	var count [3]int
//	for i := 0; i < len(nums); i++ {
//		if nums[i] < 0 || nums[i] > 2 {
//			panic("error input data")
//		}
//		count[nums[i]] += 1
//	}
//
//	setNums(0, count[0], 0, nums)
//	setNums(count[0], count[0]+count[1], 1, nums)
//	setNums(count[0]+count[1], count[0]+count[1]+count[2], 2, nums)
//}
//
//func setNums(start, end, data int, nums []int) {
//	for i := start; i < end; i++ {
//		nums[i] = data
//	}
//}

// 3 way partition
func sortColors(nums []int) {
	sortColorsImpl(nums, 0, len(nums)-1)
}

func sortColorsImpl(nums []int, l, r int) {
	zero := l - 1 // nums[l, zero] = 0
	two := r + 1  // nums[two, r] = 2

	for i := 0; i < two; {
		if nums[i] == 1 { // nums[zero+1, i-1] = 1
			i++
		} else if nums[i] == 2 {
			two--
			nums[i], nums[two] = nums[two], nums[i]
		} else {
			if nums[i] != 0 {
				panic("error")
			}
			zero++
			nums[i], nums[zero] = nums[zero], nums[i]
			i++
		}
	}
}

func main() {
	var nums = []int{0, 0, 2, 1, 0, 2}
	sortColors(nums)
	fmt.Println(nums)
}
