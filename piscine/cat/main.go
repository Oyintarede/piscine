package main

import (
	"io"
	"os"
)

func printFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, copyErr := io.Copy(os.Stdout, file)
	return copyErr
}

func main() {
	args := os.Args[1:]

	// No arguments → read from stdin
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	// Print each file in order
	for _, filename := range args {
		err := printFile(filename)
		if err != nil {
			os.Stderr.Write([]byte("ERROR: " + err.Error() + "\n"))
			os.Exit(1)
		}
	}
}
