package main

import "fmt"

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
//
//Example 1:
//[[1,3,1],
// [1,5,1],
// [4,2,1]]
//Given the above grid map, return 7. Because the path 1→3→1→1→1 minimizes the sum.

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i-1 >= 0 && j-1 >= 0 && grid[i-1][j] <= grid[i][j-1] {
				grid[i][j] += grid[i-1][j]
				continue
			}
			if i-1 >= 0 && j-1 >= 0 && grid[i-1][j] > grid[i][j-1] {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if i-1 >= 0 {
				grid[i][j] += grid[i-1][j]
			}
			if j-1 >= 0 {
				grid[i][j] += grid[i][j-1]
			}
		}
	}

	return grid[m-1][n-1]
}

func main() {
	var grid = [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}

	fmt.Println(minPathSum(grid))
}
