package main

import "fmt"

//Given a 2D board and a word, find if the word exists in the grid.
//
//The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
//For example,
//Given board =
//
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//word = "ABCCED", -> returns true,
//word = "SEE", -> returns true,
//word = "ABCB", -> returns false
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return word == ""
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if find(board, x, y, word, 0) {
				return true
			}
		}
	}
	return false
}

func find(board [][]byte, x, y int, word string, index int) bool {
	if x < 0 || x > len(board)-1 || y < 0 || y > len(board[0])-1 {
		return false
	}
	if index == len(word)-1 {
		return board[x][y] == word[index]
	}
	if word[index] == board[x][y] {
		board[x][y] ^= 255
		index++
		// up down left right
		if find(board, x-1, y, word, index) || find(board, x+1, y, word, index) || find(board, x, y-1, word, index) || find(board, x, y+1, word, index) {
			return true
		}
		board[x][y] ^= 255
		index--
	}
	return false
}
func main() {
	var board = [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
}
