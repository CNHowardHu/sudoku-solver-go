package main

import (
	"fmt"
	"os"
)

func translate(Sudoku *([9][9]byte), s string) bool {
	if len(s) != 81 {
		return false
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if s[9*i+j] == '0' || s[9*i+j] == '.' {
				Sudoku[i][j] = byte(0)
			} else if s[9*i+j] >= '1' && s[9*i+j] <= '9' {
				Sudoku[i][j] = byte(s[9*i+j] - '0')
			} else {
				return false
			}
		}
	}
	return judgeBoard(Sudoku, false)
}

func printBoard(Sudoku *([9][9]byte)) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%2d ", (*Sudoku)[i][j])
		}
		fmt.Println("")
	}
}

func pressToExit() {
	fmt.Println("Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
