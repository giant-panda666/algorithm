package main

import "fmt"

//Given two arrays, write a function to compute their intersection.
//
//Example:
//Given nums1 = [1, 2, 2, 1], nums2 = [2, 2], return [2].
//
//Note:
//Each element in the result must be unique.
//The result can be in any order.

func intersection(nums1 []int, nums2 []int) []int {
	var record = make(map[int]struct{})
	for _, v := range nums1 {
		record[v] = struct{}{}
	}

	var resultSet = make(map[int]struct{})
	for _, v := range nums2 {
		if _, ok := record[v]; ok {
			resultSet[v] = struct{}{}
		}
	}

	var res []int
	for k := range resultSet {
		res = append(res, k)
	}
	return res
}

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println("expect [2], result:", intersection(nums1, nums2))
}
