package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	counts := [10]int{}

	if n == 0 {
		z01.PrintRune('0')
		return
	}

	for n > 0 {
		digit := n % 10
		counts[digit]++
		n = n / 10
	}

	for i := 0; i <= 9; i++ {
		for counts[i] > 0 {
			z01.PrintRune(rune('0' + i))
			counts[i]--
		}
	}
}
