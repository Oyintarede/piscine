package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		u := uint(-n)
		printUInt(u)
		return
	}

	printUInt(uint(n))
}

func printUInt(u uint) {
	if u >= 10 {
		printUInt(u / 10)
	}
	z01.PrintRune(rune('0' + u%10))
}
