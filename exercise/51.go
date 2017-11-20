package main

import "fmt"

//The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.
//
//
//
//Given an integer n, return all distinct solutions to the n-queens puzzle.
//
//Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a queen and an empty space respectively.
//
//For example,
//There exist two distinct solutions to the 4-queens puzzle:
//
//[
// [".Q..",  // Solution 1
//  "...Q",
//  "Q...",
//  "..Q."],
//
// ["..Q.",  // Solution 2
//  "Q...",
//  "...Q",
//  ".Q.."]
//]

var col, dia1, dia2 []bool

func solveNQueens(n int) [][]string {
	if n <= 0 {
		return nil
	}
	col = make([]bool, n)
	dia1 = make([]bool, 2*n-1)
	dia2 = make([]bool, 2*n-1)

	var cur []string
	var res [][]string
	putQueen(n, 0, &cur, &res)
	return res
}

func putQueen(n int, index int, row *[]string, res *[][]string) {
	if n == index {
		tmp := make([]string, len(*row))
		copy(tmp, *row)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if isLegal(index, i, n) {
			col[i] = true
			dia1[index+i] = true
			dia2[i-index+n-1] = true
			tmp := ""
			for j := 0; j < n; j++ {
				if j == i {
					tmp += "Q"
				} else {
					tmp += "."
				}
			}
			*row = append(*row, tmp)
			putQueen(n, index+1, row, res)
			col[i] = false
			dia1[index+i] = false
			dia2[i-index+n-1] = false
			*row = (*row)[:len(*row)-1]
		}
	}
}

func isLegal(index, i int, n int) bool {
	return !col[i] && !dia1[index+i] && !dia2[i-index+n-1]
}

func main() {
	res := solveNQueens(8)
	for i, v := range res {
		fmt.Println(i, "Solution")
		for _, vv := range v {
			fmt.Println(vv)
		}
	}
}
