package main

import "fmt"

//Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.
//
//For example, given the following triangle
//[
//     [2],
//    [3,4],
//   [6,5,7],
//  [4,1,8,3]
//]
//The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).
//
//Note:
//Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
			} else {
				min := triangle[i-1][j-1]
				if j < len(triangle[i-1]) && min > triangle[i-1][j] {
					min = triangle[i-1][j]
				}
				triangle[i][j] += min
			}
		}
	}

	var res = triangle[len(triangle)-1][0]
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		if triangle[len(triangle)-1][i] < res {
			res = triangle[len(triangle)-1][i]
		}
	}
	return res
}

func main() {
	var triangle = [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}

	fmt.Println(minimumTotal(triangle))
}
