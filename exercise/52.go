package main

import "fmt"

//Follow up for N-Queens problem.
//
//Now, instead outputting board configurations, return the total number of distinct solutions.

var col, dia1, dia2 []bool

func totalNQueens(n int) int {
	if n <= 0 {
		return 0
	}
	col = make([]bool, n)
	dia1 = make([]bool, 2*n-1)
	dia2 = make([]bool, 2*n-1)

	var res int
	putQueen(n, 0, &res)
	return res
}

func putQueen(n int, index int, res *int) {
	if n == index {
		*res += 1
		return
	}

	for i := 0; i < n; i++ {
		if isLegal(index, i, n) {
			col[i] = true
			dia1[index+i] = true
			dia2[i-index+n-1] = true
			putQueen(n, index+1, res)
			col[i] = false
			dia1[index+i] = false
			dia2[i-index+n-1] = false
		}
	}
}

func isLegal(index, i int, n int) bool {
	return !col[i] && !dia1[index+i] && !dia2[i-index+n-1]
}

func main() {
	fmt.Println(totalNQueens(8))
}
