package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	for _, r := range args {
		if r == "01" || r == "galaxy" || r == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
