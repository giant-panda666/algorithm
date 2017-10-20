package main

import "fmt"

//Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
//
//Note:
//You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.

func merge(nums1 []int, m int, nums2 []int, n int) {
	index0, index1, index2 := m+n-1, m-1, n-1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[index0] = nums1[index1]
			index0--
			index1--
		} else {
			nums1[index0] = nums2[index2]
			index0--
			index2--
		}
	}
	for index1 >= 0 {
		nums1[index0] = nums1[index1]
		index0--
		index1--
	}
	for index2 >= 0 {
		nums1[index0] = nums2[index2]
		index0--
		index2--
	}
}

func main() {
	nums1 := []int{1, 3, 5, 8, 9, 0, 0, 0, 0}
	m := 5
	nums2 := []int{2, 4, 6, 10}
	n := 4
	merge(nums1, m, nums2, n)

	fmt.Println(nums1)
}
