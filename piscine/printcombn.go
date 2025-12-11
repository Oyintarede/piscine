package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	var comb [10]int
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	printComb(comb[:n])

	for nextComb(comb[:n]) {
		printComb(comb[:n])
	}
	z01.PrintRune('\n')
}

func printComb(c []int) {
	for _, d := range c {
		z01.PrintRune(rune('0' + d))
	}
	if !isLast(c) {
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
}

func nextComb(c []int) bool {
	n := len(c)
	for i := n - 1; i >= 0; i-- {
		if c[i] < 10-n+i {
			c[i]++
			for j := i + 1; j < n; j++ {
				c[j] = c[j-1] + 1
			}
			return true
		}
	}
	return false
}

func isLast(c []int) bool {
	n := len(c)
	for i := 0; i < n; i++ {
		if c[i] != 10-n+i {
			return false
		}
	}
	return true
}
