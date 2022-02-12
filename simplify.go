package main

func simplify(Sudoku *([9][9]byte)) bool {
	methods := [...]func(Sudoku *([9][9]byte)) bool{
		simplifyLastValue,
		simplifyHiddenSingleInBox,
		simplifyHiddenSingleInRowColumn,
		simplifyNakedSingle}
	for endLoop := false; !endLoop; {
		endLoop = true
		for _, m := range methods {
			for p := true; p; {
				p = m(Sudoku)
				endLoop = endLoop && !p
			}
		}
	}
	return judgeBoard(Sudoku, false)
}

func simplifyLastValue(Sudoku *([9][9]byte)) bool {
	res := false
	for i := 0; i < 9; i++ {
		var cnt, p [3]int
		var q [3][9]bool
		for j := 0; j < 9; j++ {
			for k := 0; k < 3; k++ {
				x, y := ijk2xy(i, j, k)
				if Sudoku[x][y] == 0 {
					cnt[k]++
					p[k] = j
				} else {
					if q[k][Sudoku[x][y]-1] {
						return false
					} else {
						q[k][Sudoku[x][y]-1] = true
					}
				}
			}
		}
		for k := 0; k < 3; k++ {
			if cnt[k] == 1 {
				x, y := ijk2xy(i, p[k], k)
				Sudoku[x][y] = func() byte {
					for num := byte(1); num <= 9; num++ {
						if !q[k][num-1] {
							return num
						}
					}
					return 0
				}()
				res = true
			}
		}
	}
	return res && judgeBoard(Sudoku, false)
}

func simplifyHiddenSingleInBox(Sudoku *([9][9]byte)) bool {
	res := false
	for i := 0; i < 9; i++ {
		var p [9]bool
		for j := 0; j < 9; j++ {
			if num := ijk2num(Sudoku, i, j, 2); num != 0 {
				if p[num-1] {
					return false
				} else {
					p[num-1] = true
				}
			}
		}
		for num := byte(1); num <= 9; num++ {
			if !p[num-1] {
				var cnt, pos int
				var q [9]bool
				for j := 0; j < 9; j++ {
					x, y := ijk2xy(i, j, 2)
					if Sudoku[x][y] == 0 {
						q[j] = true
						for l := 0; l < 9; l++ {
							if Sudoku[x][l] == num || Sudoku[l][y] == num {
								q[j] = false
								break
							}
						}
					}
					if q[j] {
						cnt++
						pos = j
					}
				}
				if cnt == 1 {
					x, y := ijk2xy(i, pos, 2)
					Sudoku[x][y] = num
					res = true
				}
			}
		}
	}
	return res && judgeBoard(Sudoku, false)
}

func simplifyHiddenSingleInRowColumn(Sudoku *([9][9]byte)) bool {
	res := false
	for k := 0; k <= 1; k++ {
		for i := 0; i < 9; i++ {
			var p [9]bool
			for j := 0; j < 9; j++ {
				if num := ijk2num(Sudoku, i, j, k); num != 0 {
					if p[num-1] {
						return false
					} else {
						p[num-1] = true
					}
				}
			}
			for num := byte(1); num <= 9; num++ {
				if !p[num-1] {
					var cnt, pos int
					var q [9]bool
					for j := 0; j < 9; j++ {
						x, y := ijk2xy(i, j, k)
						if Sudoku[x][y] == 0 {
							q[j] = true
							irc, _ := xyk2ij(x, y, 1-k)
							ibox, _ := xyk2ij(x, y, 2)
							for l := 0; l < 9; l++ {
								if ijk2num(Sudoku, irc, l, 1-k) == num ||
									ijk2num(Sudoku, ibox, l, 2) == num {
									q[j] = false
									break
								}
							}
						}
						if q[j] {
							cnt++
							pos = j
						}
					}
					if cnt == 1 {
						x, y := ijk2xy(i, pos, k)
						Sudoku[x][y] = num
						res = true
					}
				}
			}
		}
	}
	return res && judgeBoard(Sudoku, false)
}

func simplifyNakedSingle(Sudoku *([9][9]byte)) bool {
	res := false
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			if Sudoku[x][y] == 0 {
				var p [9]bool
				for k := 0; k < 3; k++ {
					i, _ := xyk2ij(x, y, k)
					for l := 0; l < 9; l++ {
						if num := ijk2num(Sudoku, i, l, k); num != 0 {
							p[num-1] = true
						}
					}
				}
				var cnt, pos int
				for num := 1; num <= 9; num++ {
					if !p[num-1] {
						cnt++
						pos = num
					}
				}
				if cnt == 1 {
					Sudoku[x][y] = byte(pos)
					res = true
				}
			}
		}
	}
	return res && judgeBoard(Sudoku, false)
}

func judgeBoard(Sudoku *([9][9]byte), needComplete bool) bool {
	for i := 0; i < 9; i++ {
		var p [9][3]bool
		for j := 0; j < 9; j++ {
			for k := 0; k < 3; k++ {
				x, y := ijk2xy(i, j, k)
				if Sudoku[x][y] != 0 {
					if p[Sudoku[x][y]-1][k] {
						return false
					}
					p[Sudoku[x][y]-1][k] = true
				}
			}
		}
		if needComplete {
			for num := 1; num <= 9; num++ {
				for k := 0; k < 2; k++ {
					if !p[num-1][k] {
						return false
					}
				}
			}
		}
	}
	return true
}

func ijk2num(Sudoku *([9][9]byte), i, j, k int) byte {
	x, y := ijk2xy(i, j, k)
	return Sudoku[x][y]
}

func ijk2xy(i, j, k int) (int, int) {
	switch k {
	case 0:
		return i, j
	case 1:
		return j, i
	case 2:
		return (i/3*3 + j/3), (i%3*3 + j%3)
	default:
		return i, j
	}
}

func xyk2ij(x, y, k int) (int, int) {
	switch k {
	case 0:
		return x, y
	case 1:
		return y, x
	case 2:
		return (x/3*3 + y/3), (x%3*3 + y%3)
	default:
		return x, y
	}
}
