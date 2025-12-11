package main

import (
	"fmt"
	"os"
)

func printTail(filename string, n int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	size := stat.Size()
	start := int64(0)
	if int64(n) < size {
		start = size - int64(n)
	}

	_, err = file.Seek(start, 0)
	if err != nil {
		return err
	}

	buf := make([]byte, size-start)
	_, err = file.Read(buf)
	if err != nil && err.Error() != "EOF" {
		return err
	}

	fmt.Printf("%s", buf)
	return nil
}

func main() {
	args := os.Args[1:]
	hadError := false

	if len(args) < 3 || args[0] != "-c" {
		fmt.Printf("Usage: ./program -c <number> <file1> [file2 ...]\n")
		os.Exit(1)
	}

	var n int
	_, err := fmt.Sscanf(args[1], "%d", &n)
	if err != nil || n <= 0 {
		fmt.Printf("Invalid number of bytes\n")
		os.Exit(1)
	}

	files := args[2:]

	for i, f := range files {
		// Check if the file exists first
		file, err := os.Open(f)
		if err != nil {
			fmt.Printf("%s\n", err)
			hadError = true
			continue
		}
		file.Close()

		// Print header only for existing files when multiple files are given
		if len(files) > 1 {
			if i > 0 {
				fmt.Printf("\n")
			}
			fmt.Printf("==> %s <==\n", f)
		}

		err = printTail(f, n)
		if err != nil {
			fmt.Printf("%s\n", err)
			hadError = true
		}
	}

	if hadError {
		os.Exit(1)
	}
}
