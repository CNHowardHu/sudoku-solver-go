package main

import (
	"fmt"
)

// valid input format:

// 980700600700000080006050000400003002007940060000000400010000003009500070000020100
// 98.7..6..7......8...6.5....4....3..2..794..6.......4...1......3..95...7.....2.1..

func main() {
	for true {
		var s string
		var Sudoku [9][9]byte
		fmt.Scanln(&s)
		if translate(&Sudoku, s) {
			if DFS_ng(&Sudoku, 0) {
				printBoard(&Sudoku)
			} else {
				fmt.Println("Failed to solve this!")
			}
		} else {
			fmt.Println("Invalid task!")
			break
		}
	}
	pressToExit()
}

// https://leetcode-cn.com/problems/sudoku-solver/

// func solveSudoku(board [][]byte) {
// 	var Sudoku = [9][9]byte{}
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] == '.' {
// 				Sudoku[i][j] = byte(0)
// 			} else {
// 				Sudoku[i][j] = byte(board[i][j] - '0')
// 			}
// 		}
// 	}
// 	DFS_ng(&Sudoku, 0)
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			board[i][j] = byte(Sudoku[i][j] + '0')
// 		}
// 	}
// }
