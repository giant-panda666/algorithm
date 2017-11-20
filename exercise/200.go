package main

import "fmt"

//Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
//Example 1:
//
//11110
//11010
//11000
//00000
//Answer: 1
//
//Example 2:
//
//11000
//11000
//00100
//00011
//Answer: 3

func numIslands(grid [][]byte) int {
	if grid == nil {
		return 0
	}
	var mark = make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		mark[i] = make([]bool, len(grid[i]))
	}

	var count = 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 && !mark[x][y] {
				count++
				backTrace(grid, x, y, mark)
			}
		}
	}
	return count
}

func backTrace(grid [][]byte, x, y int, mark [][]bool) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}

	if grid[x][y] == 1 && !mark[x][y] {
		mark[x][y] = true
		backTrace(grid, x, y+1, mark)
		backTrace(grid, x+1, y, mark)
		backTrace(grid, x, y-1, mark)
		backTrace(grid, x-1, y, mark)
	}
}

func main() {
	var grid = [][]byte{
		[]byte{1, 1, 0, 0, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 1, 0, 0},
		[]byte{0, 0, 0, 1, 1},
	}
	fmt.Println(numIslands(grid))

	var grid2 = [][]byte{
		[]byte{1, 1, 1, 1, 0},
		[]byte{1, 1, 0, 1, 0},
		[]byte{1, 1, 0, 0, 0},
		[]byte{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(grid2))
}
