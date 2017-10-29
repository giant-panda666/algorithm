package main

import "fmt"

//Given two arrays, write a function to compute their intersection.
//
//Example:
//Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2, 2].
//
//Note:
//Each element in the result should appear as many times as it shows in both arrays.
//The result can be in any order.
//Follow up:
//What if the given array is already sorted? How would you optimize your algorithm?
//What if nums1's size is small compared to nums2's size? Which algorithm is better?
//What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?

func intersect(nums1 []int, nums2 []int) []int {
	var record = make(map[int]int)
	for _, v := range nums1 {
		record[v]++
	}

	var resultRecord = make(map[int]int)
	for _, v := range nums2 {
		if vv, ok := record[v]; ok && vv > 0 {
			resultRecord[v]++
			record[v]--
		}
	}

	var res []int
	for k, v := range resultRecord {
		for i := 0; i < v; i++ {
			res = append(res, k)
		}
	}
	return res
}

func main() {
	var nums1 = []int{1, 2, 2, 1}
	var nums2 = []int{2, 2}
	fmt.Println("expect [2, 2], reuslt:", intersect(nums1, nums2))
}
