package main

import "fmt"

//Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.
//
//Note: You may not slant the container and n is at least 2.

// maxArea = (j-i)*min(ai, aj), 1<=i<j<=n
// ai, aj中的较小者决定了当前面积的上界，只有将较小者往中间靠拢，找到一个更高的柱子，面积才有可能增加
func maxArea(height []int) int {
	if len(height) < 2 {
		panic("error input")
	}
	i, j := 0, len(height)-1
	max := area(i, j, height[i], height[j])
	for i < j {
		if height[i] < height[j] {
			index := i + 1
			for index < j && height[index] < height[i] {
				index++
			}
			i = index
		} else {
			index := j - 1
			for index > i && height[index] < height[j] {
				index--
			}
			j = index
		}
		if area(i, j, height[i], height[j]) > max {
			max = area(i, j, height[i], height[j])
		}
		if i+1 == j {
			return max
		}
	}
	return max
}

// assert j > i
func area(i, j, hi, hj int) int {
	return (j - i) * min(hi, hj)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var height = []int{1, 2, 3}
	fmt.Println("expect 2", maxArea(height))
}
