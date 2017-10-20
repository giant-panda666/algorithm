package main

import (
	"fmt"
	"math/rand"
)

//Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
//For example,
//Given [3,2,1,5,6,4] and k = 2, return 5.
//
//Note:
//You may assume k is always valid, 1 ≤ k ≤ array's length.

func findKthLargest(nums []int, k int) int {
	return findKthLargestImpl(nums, 0, len(nums)-1, k-1)
}

func findKthLargestImpl(nums []int, l, r, k int) int {
	if l > r || k < 0 || k >= len(nums) {
		panic("error")
	}
	// 随机选择一个值作为partition的分界点
	privot := rand.Int()%(r-l+1) + l
	val := nums[privot]

	i := l - 1 // nums[l, ... i] 比val小
	j := r + 1 // nums[j, ... r] 比val大
	index := l // nums[i+1, ... index-1] 等于val

	for index < j {
		if nums[index] == val {
			index++
		} else if nums[index] > val {
			i++
			nums[index], nums[i] = nums[i], nums[index]
			index++
		} else if nums[index] < val {
			j--
			nums[index], nums[j] = nums[j], nums[index]
		}
	}
	nums[index-1] = val

	if i+1 <= k && index-1 >= k {
		return val
	}
	if index-1 < k {
		return findKthLargestImpl(nums, j, r, k)
	}
	return findKthLargestImpl(nums, l, i, k)
}

func main() {
	var nums = []int{3, 2, 1, 5, 6, 4}
	fmt.Println("expect 5, returned:", findKthLargest(nums, 2))
}
