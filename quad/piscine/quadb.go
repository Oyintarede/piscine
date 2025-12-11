package piscine

import "fmt"

func QuadB(x,y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 1; row <=y; row++ {
		for col :=1; col <= x; col++ {

			if (row == 1 && col == 1) || (row == y && col == x ) {
				fmt.Print("/")
			} else if (row == 1 && col == x) || (row == y && col == 1) {
				fmt.Print("\\")
			} else if (row == 1 || row == y) || (col == 1 || col == x) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
