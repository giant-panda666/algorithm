package main

//Follow up for "Unique Paths":
//
//Now consider if some obstacles are added to the grids. How many unique paths would there be?
//
//An obstacle and empty space is marked as 1 and 0 respectively in the grid.
//
//For example,
//There is one obstacle in the middle of a 3x3 grid as illustrated below.
//
//[
//  [0,0,0],
//  [0,1,0],
//  [0,0,0]
//]
//The total number of unique paths is 2.
//
//Note: m and n will be at most 100

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) < 1 || obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	var graph = make([][]int, m)
	for i := 0; i < m; i++ {
		graph[i] = make([]int, n)
	}

	if obstacleGrid[m-1][n-1] == 1 {
		graph[m-1][n-1] = 0
	} else {
		graph[m-1][n-1] = 1
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j+1 <= n-1 && obstacleGrid[i][j+1] != 1 {
				graph[i][j] += graph[i][j+1]
			}
			if i+1 <= m-1 && obstacleGrid[i+1][j] != 1 {
				graph[i][j] += graph[i+1][j]
			}
		}
	}

	return graph[0][0]
}
