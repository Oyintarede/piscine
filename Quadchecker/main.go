package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	input := string(data)

	input = strings.TrimRight(input, "\n")
	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0])

	quadList := map[string]string{
		"quadA": QuadA(w, h),
		"quadB": QuadB(w, h),
		"quadC": QuadC(w, h),
		"quadD": QuadD(w, h),
		"quadE": QuadE(w, h),
	}

	matches := []string{}

	for name, pattern := range quadList {
		pattern = strings.TrimRight(pattern, "\n")
		if pattern == input {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", name, w, h))
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	order := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	final := []string{}

	for _, q := range order {
		for _, m := range matches {
			if strings.Contains(m, q) {
				final = append(final, m)
			}
		}
	}

	fmt.Println(strings.Join(final, " || "))
}

func QuadA(x, y int) string {
	var b strings.Builder
	if x <= 0 || y <= 0 {
		return ""
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				b.WriteRune('o')
			} else if i == 0 && j == x-1 {
				b.WriteRune('o')
			} else if i == y-1 && j == 0 {
				b.WriteRune('o')
			} else if i == y-1 && j == x-1 {
				b.WriteRune('o')
			} else if i == 0 || i == y-1 {
				b.WriteRune('-')
			} else if j == 0 || j == x-1 {
				b.WriteRune('|')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func QuadB(x, y int) string {
	var b strings.Builder
	if x <= 0 || y <= 0 {
		return ""
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				b.WriteRune('/')
			} else if i == 0 && j == x-1 {
				b.WriteRune('\\')
			} else if i == y-1 && j == 0 {
				b.WriteRune('\\')
			} else if i == y-1 && j == x-1 {
				b.WriteRune('/')
			} else if i == 0 || i == y-1 {
				b.WriteRune('*')
			} else if j == 0 || j == x-1 {
				b.WriteRune('*')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func QuadC(x, y int) string {
	var b strings.Builder
	if x <= 0 || y <= 0 {
		return ""
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				b.WriteRune('A')
			} else if i == 0 && j == x-1 {
				b.WriteRune('A')
			} else if i == y-1 && j == 0 {
				b.WriteRune('C')
			} else if i == y-1 && j == x-1 {
				b.WriteRune('C')
			} else if i == 0 || i == y-1 {
				b.WriteRune('B')
			} else if j == 0 || j == x-1 {
				b.WriteRune('B')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func QuadD(x, y int) string {
	var b strings.Builder
	if x <= 0 || y <= 0 {
		return ""
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				b.WriteRune('A')
			} else if i == 0 && j == x-1 {
				b.WriteRune('C')
			} else if i == y-1 && j == 0 {
				b.WriteRune('A')
			} else if i == y-1 && j == x-1 {
				b.WriteRune('C')
			} else if i == 0 || i == y-1 {
				b.WriteRune('B')
			} else if j == 0 || j == x-1 {
				b.WriteRune('B')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func QuadE(x, y int) string {
	var b strings.Builder

	if x <= 0 || y <= 0 {
		return ""
	}

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i == 0 && j == 0 {
				b.WriteRune('A')
			} else if i == 0 && j == x-1 {
				b.WriteRune('C')
			} else if i == y-1 && j == 0 {
				b.WriteRune('C')
			} else if i == y-1 && j == x-1 {
				b.WriteRune('A')
			} else if i == 0 || i == y-1 {
				b.WriteRune('B')
			} else if j == 0 || j == x-1 {
				b.WriteRune('B')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}
