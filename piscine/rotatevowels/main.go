package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printRune(r rune) {
	z01.PrintRune(r)
}

func printStrSafe(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func newLine() {
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		newLine()
		return
	}

	var runes []rune
	for i, arg := range args {
		for _, r := range arg {
			runes = append(runes, r)
		}
		if i < len(args)-1 {
			runes = append(runes, ' ')
		}
	}

	var vowelPos []int
	for i, r := range runes {
		if isVowel(r) {
			vowelPos = append(vowelPos, i)
		}
	}

	if len(vowelPos) == 0 {
		for _, r := range runes {
			printRune(r)
		}
		newLine()
		return
	}

	left := 0
	right := len(vowelPos) - 1

	for left < right {
		i := vowelPos[left]
		j := vowelPos[right]

		runes[i], runes[j] = runes[j], runes[i]

		left++
		right--
	}

	for _, r := range runes {
		printRune(r)
	}
	newLine()
}
