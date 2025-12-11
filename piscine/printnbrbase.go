package piscine

import "github.com/01-edu/z01"

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func printNbrBaseRec(n uint, base string) {
	baseLen := uint(len(base))
	if n >= baseLen {
		printNbrBaseRec(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		// convert to uint to safely handle MinInt
		printNbrBaseRec(uint(-nbr), base)
	} else {
		printNbrBaseRec(uint(nbr), base)
	}
}
