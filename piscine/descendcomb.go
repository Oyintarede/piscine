package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			ii := i
			jj := j
			if j == '0' {
				ii--
				jj = '9'
			} else {
				jj--
			}
			for k := ii; k >= '0'; k-- {
				start := '9'
				if k == ii {
					start = jj
				}
				for l := start; l >= '0'; l-- {
					if !first {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					z01.PrintRune(l)
					first = false
				}
			}
		}
	}
}
