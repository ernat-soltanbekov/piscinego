package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(0, &board)
}

func solve(row int, board *[8]int) {
	if row == 8 {
		printBoard(board)
		return
	}
	for col := 0; col < 8; col++ {
		if isSafe(row, col, board) {
			board[row] = col
			solve(row+1, board)
		}
	}
}

func isSafe(row, col int, board *[8]int) bool {
	for i := 0; i < row; i++ {
		prevCol := board[i]
		if prevCol == col || prevCol-col == i-row || prevCol-col == row-i {
			return false
		}
	}
	return true
}

func printBoard(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '1'))
	}
	z01.PrintRune('\n')
}
