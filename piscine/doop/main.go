package main

import (
	"os"
)

func Atoi2(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign := 1
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	if start >= len(s) {
		return 0, false
	}
	n := 0
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		if n > (1<<31-1)/10 {
			return 0, false
		}
		n = n * 10
		digit := int(s[i] - '0')
		if n > (1<<31-1)-digit {
			return 0, false
		}
		n = n + digit
	}
	result := n * sign
	if sign == 1 && result < 0 {
		return 0, false
	}
	if sign == -1 && result > 0 {
		return 0, false
	}
	return result, true
}

func PrintNbr2(n int) {
	if n == 0 {
		os.Stdout.WriteString("0")
		return
	}
	if n == -1<<31 {
		os.Stdout.WriteString("-2147483648")
		return
	}
	if n < 0 {
		os.Stdout.WriteString("-")
		n = -n
	}
	var digits []byte
	for n > 0 {
		digits = append([]byte{byte(n%10 + '0')}, digits...)
		n /= 10
	}
	os.Stdout.Write(digits)
}

func PrintStr2(s string) {
	os.Stdout.WriteString(s)
	os.Stdout.WriteString("\n")
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	aStr, operator, bStr := os.Args[1], os.Args[2], os.Args[3]
	a, ok1 := Atoi2(aStr)
	b, ok2 := Atoi2(bStr)
	if !ok1 || !ok2 {
		return
	}
	var result int
	valid := true
	switch operator {
	case "+":
		if (b > 0 && a > (1<<31-1)-b) || (b < 0 && a < (-1<<31)-b) {
			return
		}
		result = a + b
	case "-":
		if (b < 0 && a > (1<<31-1)+b) || (b > 0 && a < (-1<<31)+b) {
			return
		}
		result = a - b
	case "*":
		if a > 0 {
			if b > 0 && a > (1<<31-1)/b {
				return
			} else if b < 0 && b < (-1<<31)/a {
				return
			}
		} else if a < 0 {
			if b > 0 && a < (-1<<31)/b {
				return
			} else if b < 0 && a < (1<<31-1)/b {
				return
			}
		}
		result = a * b
	case "/":
		if b == 0 {
			PrintStr2("No division by 0")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			PrintStr2("No modulo by 0")
			return
		}
		result = a % b
	default:
		valid = false
	}
	if !valid {
		return
	}
	PrintNbr2(result)
	os.Stdout.WriteString("\n")
}
