package main

func DFS_ng(Sudoku *([9][9]byte), n int) bool {
	if n > 80 {
		return judgeBoard(Sudoku, true)
	}
	next := func() int {
		res := n + 1
		for res < 81 && Sudoku[res/9][res%9] != 0 {
			res++
		}
		return res
	}
	x, y := n/9, n%9
	if Sudoku[x][y] != 0 {
		return DFS_ng(Sudoku, next())
	} else {
		tmp := *Sudoku
		for num := byte(1); num <= 9; num++ {
			check := func() bool {
				for k := 0; k < 3; k++ {
					i, _ := xyk2ij(x, y, k)
					for j := 0; j < 9; j++ {
						if ijk2num(Sudoku, i, j, k) == num {
							return false
						}
					}
				}
				return true
			}
			if check() {
				Sudoku[x][y] = num
				if simplify(Sudoku) && DFS_ng(Sudoku, next()) {
					return true
				}
				*Sudoku = tmp
			}
		}
		return false
	}
}
