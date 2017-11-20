package main

import "fmt"

//Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.
//
//A region is captured by flipping all 'O's into 'X's in that surrounded region.
//
//For example,
//X X X X
//X O O X
//X X O X
//X O X X
//After running your function, the board should be:
//
//X X X X
//X X X X
//X X X X
//X O X X

func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])
	for i := 0; i < row; i++ {
		check(board, i, 0, row, col)
		if col > 1 {
			check(board, i, col-1, row, col)
		}
	}

	for j := 1; j < col-1; j++ {
		check(board, 0, j, row, col)
		if row > 1 {
			check(board, row-1, j, row, col)
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

func check(board [][]byte, i, j, row, col int) {
	if board[i][j] == 'O' {
		board[i][j] = '1'
		if i-1 > 0 {
			check(board, i-1, j, row, col)
		}
		if i+1 < row {
			check(board, i+1, j, row, col)
		}
		if j-1 > 0 {
			check(board, i, j-1, row, col)
		}
		if j+1 < col {
			check(board, i, j+1, row, col)
		}
	}
}

func main() {
	var board = [][]byte{
		[]byte{'X', 'O', 'X', 'X'},
		[]byte{'O', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X'},
		[]byte{'X', 'O', 'X', 'O'},
		[]byte{'O', 'X', 'O', 'X'},
	}
	solve(board)
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			fmt.Printf("%c", board[x][y])
		}
		fmt.Println()
	}
}
