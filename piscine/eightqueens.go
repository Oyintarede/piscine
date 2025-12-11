package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	backtrack(&board, 0)
}

func backtrack(board *[8]int, col int) {
	if col == 8 {
		printSolution(board)
		return
	}
	for row := 1; row <= 8; row++ {
		if safe(board, col, row) {
			board[col] = row
			backtrack(board, col+1)
		}
	}
}

func safe(board *[8]int, col int, row int) bool {
	for i := 0; i < col; i++ {
		if board[i] == row ||
			board[i]-row == i-col ||
			row-board[i] == i-col {
			return false
		}
	}
	return true
}

func printSolution(board *[8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + '0'))
	}
	z01.PrintRune('\n')
}
