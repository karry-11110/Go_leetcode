package main

import "strings"

func isValid(board [][]string, row, col int) bool {
	//判断再填充（row,col）之前是否在行列上已经有填充了
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" {
			return false
		}
	}
	for j := 0; j < col; j++ {
		if board[row][j] == "Q" {
			return false
		}
	}
	//判断在（row,col）所在对角线上是否有Q
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}

	for i, j := row, col; i >= 0 && j < len(board); i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
func solveNQueens(n int) [][]string {
	result := [][]string{}
	//制作n*n的棋盘格
	board := make([][]string, n)
	for i := 0; i < n; i++ {
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	var backtracking func(board [][]string, row int)
	backtracking = func(board [][]string, row int) {
		if row == len(board) { //棋盘是一行行来填充完成的
			tmp := make([]string, len(board))
			for i := 0; i < len(board); i++ {
				tmp[i] = strings.Join(board[i], "")
			}
			result = append(result, tmp) //!!!
		}
		for col := 0; col < len(board); col++ { //棋盘的宽度为循环的长度，即多少列
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = "Q"
			backtracking(board, row+1)
			board[row][col] = "."
		}
	}
	backtracking(board, 0)
	return result
}
