package main

import (
	"os"

	"github.com/01-edu/z01"
)

// Print a string rune by rune
func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// ASCII bubble sort
func sortASCII(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

// Check if str starts with prefix
func startsWith(str string, prefix string) bool {
	if len(str) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if str[i] != prefix[i] {
			return false
		}
	}
	return true
}

// Remove prefix manually
func removePrefix(str string, prefix string) string {
	return str[len(prefix):]
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"--help"}
	}

	insertions := ""
	order := false
	strArg := ""
	showHelp := false

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			showHelp = true
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if startsWith(arg, "--insert=") {
			insertions += removePrefix(arg, "--insert=")
		} else if startsWith(arg, "-i=") {
			insertions += removePrefix(arg, "-i=")
		} else if strArg == "" {
			strArg = arg
		}
	}

	if showHelp {
		printString("--insert")
		printString("  -i")
		printString("\t This flag inserts the string into the string passed as argument.")
		printString("--order")
		printString("  -o")
		printString("\t This flag will behave like a boolean, if it is called it will order the argument.")
		return
	}

	// Append all insertions after main argument
	if insertions != "" {
		strArg += insertions
	}

	// Order string if flag is set
	if order {
		strArg = sortASCII(strArg)
	}

	printString(strArg)
}
