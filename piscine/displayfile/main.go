package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	// No file provided
	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}

	// Too many arguments
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	// Try to open the file
	file, err := os.Open(args[0])
	if err != nil {
		// If file cannot be opened, do nothing special (default behavior not specified)
		return
	}
	defer file.Close()

	// Copy file content to stdout
	io.Copy(os.Stdout, file)
}
