package main

//A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
//
//The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
//
//How many possible unique paths are there?
//
//
//Above is a 3 x 7 grid. How many possible unique paths are there?
//
//Note: m and n will be at most 100.

func uniquePaths(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	var graph = make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
	}

	graph[m-1][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j+1 <= n-1 && i+1 <= m-1 {
				graph[i][j] = graph[i][j+1] + graph[i+1][j]
				continue
			}
			if j+1 <= n-1 {
				graph[i][j] = graph[i][j+1]
			}
			if i+1 <= m-1 {
				graph[i][j] = graph[i+1][j]
			}
		}
	}

	return graph[0][0]
}
