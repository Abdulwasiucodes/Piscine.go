package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	var solve func(col int)

	solve = func(col int) {
		if col == 8 {
			for i := 0; i < 8; i++ {
				z01.PrintRune(rune(board[i] + '1'))
			}
			z01.PrintRune('\n')
			return
		}

		for row := 0; row < 8; row++ {
			if isSafe(board, col, row) {
				board[col] = row
				solve(col + 1)
			}
		}
	}

	solve(0)
}

func isSafe(board [8]int, col, row int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row || board[i]-row == col-i || row-board[i] == col-i {
			return false
		}
	}
	return true
}
