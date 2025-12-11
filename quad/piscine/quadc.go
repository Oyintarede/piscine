package piscine

import "fmt"

func QuadC(x,y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row :=1; row <= y; row++ {
		for col :=1; col <= x; col++ {
		
			if row == 1 && col == 1 || row == 1 && col == x {
				fmt.Print("A")
			} else if row == y && col == 1 || row == y && col == x {
				fmt.Print("C")
			} else if row == 1 || row == y || col == 1 || col == x {
				fmt.Print("B")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()

	}
}
