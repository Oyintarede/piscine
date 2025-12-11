package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	// Build numbers using allowed arithmetic
	ptr.x = (10 * 4) + (1 + 1)
	ptr.y = (10 * 2) + 1
}

func main() {
	points := &point{}
	setPoint(points)

	zero := '0'
	one := zero + 1
	two := one + 1
	// correct way to get '4': two + 1 + 1  -> 50 + 1 + 1 = 52 ('4')
	four := two + (one - zero) + (one - zero)

	out := []rune{
		'x', ' ', '=', ' ',
		four, two,
		',', ' ', 'y', ' ', '=', ' ',
		two, one,
		'\n',
	}

	for _, r := range out {
		z01.PrintRune(r)
	}

	_ = points
}
