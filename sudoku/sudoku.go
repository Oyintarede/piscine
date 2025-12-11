package main

import (
	"fmt"
	"os"
)

var field = [9][9]int{}

func draw() {
	for _, row := range field {
		for i, val := range row {
			fmt.Printf("%v", val)

			if i < len(row)-1 {
				fmt.Printf(" ")
			}
		}

		fmt.Println()
	}
}

func canPut(x int, y int, value int) bool {
	return !alreadyInVertical(x, y, value) &&
		!alreadyInHorizontal(x, y, value) &&
		!alreadyInSquare(x, y, value)
}

func alreadyInVertical(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[i][x] == value {
			return true
		}
	}
	return false
}

func alreadyInHorizontal(x int, y int, value int) bool {
	for i := range [9]int{} {
		if field[y][i] == value {
			return true
		}
	}
	return false
}

func alreadyInSquare(x int, y int, value int) bool {
	sx, sy := int(x/3)*3, int(y/3)*3
	for dy := range [3]int{} {
		for dx := range [3]int{} {
			if field[sy+dy][sx+dx] == value {
				return true
			}
		}
	}
	return false
}

func next(x int, y int) (int, int) {
	nextX, nextY := (x+1)%9, y
	if nextX == 0 {
		nextY = y + 1
	}
	return nextX, nextY
}

func solve(x int, y int) bool {
	if y == 9 {
		return true
	}
	if field[y][x] != 0 {
		return solve(next(x, y))
	} else {
		for i := range [9]int{} {
			var v = i + 1
			if canPut(x, y, v) {
				field[y][x] = v
				if solve(next(x, y)) {
					return true
				}
				field[y][x] = 0
			}
		}
		return false
	}
}

func main() {
	rows := os.Args
	if len(rows) != 10 {
		fmt.Println("Error")
		return
	}

	for i := 0; i < 9; i++ {
		newRows := rows[i+1]
		if len(newRows) != 9 {
			fmt.Println("Error")
			return
		}
		for j, char := range newRows {
			if char == '.' {
				field[i][j] = 0
			} else {
				if char >= '1' && char <= '9' {
					digitValue := int(char - '0')
					field[i][j] = digitValue
				} else {
					fmt.Println("Error")
					return
				}

			}
		}
	}

	if solve(0, 0) {
		draw()
	} else {
		fmt.Println("Error")
	}
}
