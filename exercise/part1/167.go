package main

import "fmt"

//Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.
//
//The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.
//
//You may assume that each input would have exactly one solution and you may not use the same element twice.
//
//Input: numbers={2, 7, 11, 15}, target=9
//Output: index1=1, index2=2

// 二分查找
func twoSum(numbers []int, target int) []int {
	var ret []int
	for i := 0; i < len(numbers); i++ {
		l := i + 1
		r := len(numbers) - 1
		mid := l + (r-l)/2
		for r >= l {
			sum := numbers[mid] + numbers[i]
			if sum == target {
				ret = append(ret, i+1, mid+1)
				return ret
			} else if sum > target {
				r = mid - 1
				mid = l + (r-l)/2
			} else if sum < target {
				l = mid + 1
				mid = l + (r-l)/2
			}

		}
	}
	return nil
}

// 一前一后的缩近
func twoSum2(numbers []int, target int) []int {
	var ret = []int{0, 0}
	i, j := 0, len(numbers)-1
	for j > i {
		sum := numbers[i] + numbers[j]
		if sum == target {
			ret[0], ret[1] = i+1, j+1
			return ret
		} else if sum > target {
			j--
		} else {
			i++
		}
	}

	return nil
}

func main() {
	var numbers = []int{1, 2, 3, 4, 4, 9, 56, 90}
	fmt.Println(twoSum(numbers, 8))
	fmt.Println(twoSum2(numbers, 8))
}
