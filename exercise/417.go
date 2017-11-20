package main

import "fmt"

//Given an m x n matrix of non-negative integers representing the height of each unit cell in a continent, the "Pacific ocean" touches the left and top edges of the matrix and the "Atlantic ocean" touches the right and bottom edges.
//
//Water can only flow in four directions (up, down, left, or right) from a cell to another one with height equal or lower.
//
//Find the list of grid coordinates where water can flow to both the Pacific and Atlantic ocean.
//
//Note:
//The order of returned grid coordinates does not matter.
//Both m and n are less than 150.
//Example:
//
//Given the following 5x5 matrix:
//
//  Pacific ~   ~   ~   ~   ~
//       ~  1   2   2   3  (5) *
//       ~  3   2   3  (4) (4) *
//       ~  2   4  (5)  3   1  *
//       ~ (6) (7)  1   4   5  *
//       ~ (5)  1   1   2   4  *
//          *   *   *   *   * Atlantic
//
//Return:
//
//[[0, 4], [1, 3], [1, 4], [2, 2], [3, 0], [3, 1], [4, 0]] (positions with parentheses in above matrix).

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return nil
	}
	m, n := len(matrix), len(matrix[0])
	var top, toa = make([][]bool, m), make([][]bool, m)
	var pqueue, aqueue [][]int
	for i := 0; i < m; i++ {
		top[i] = make([]bool, n)
		toa[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		top[i][0] = true
		pqueue = append(pqueue, []int{i, 0})
		toa[i][n-1] = true
		aqueue = append(aqueue, []int{i, n - 1})
	}
	for i := 0; i < n; i++ {
		top[0][i] = true
		pqueue = append(pqueue, []int{0, i})
		toa[m-1][i] = true
		aqueue = append(aqueue, []int{m - 1, i})
	}

	var dir = [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}}
	bfs := func(to [][]bool, queue [][]int) {
		for len(queue) > 0 {
			i, j := queue[0][0], queue[0][1]
			queue = queue[1:]
			for _, v := range dir {
				newi, newj := i+v[0], j+v[1]
				if newi >= 0 && newi < m && newj >= 0 && newj < n && !to[newi][newj] && matrix[newi][newj] >= matrix[i][j] {
					to[newi][newj] = true
					queue = append(queue, []int{newi, newj})
				}
			}
		}
	}

	bfs(toa, aqueue)
	bfs(top, pqueue)
	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if toa[i][j] && top[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func main() {
	//	var matrix = [][]int{
	//		[]int{1, 2, 2, 3, 5},
	//		[]int{3, 2, 3, 4, 4},
	//		[]int{2, 4, 5, 3, 1},
	//		[]int{6, 7, 1, 4, 5},
	//		[]int{5, 1, 1, 2, 4},
	//	}
	//
	//	fmt.Println(pacificAtlantic(matrix))

	var m2 = [][]int{
		[]int{1, 1},
		[]int{1, 1},
		[]int{1, 1},
	}
	fmt.Println(m2[0])
	fmt.Println(m2[1])
	fmt.Println(m2[2])
	fmt.Println(pacificAtlantic(m2))
}
