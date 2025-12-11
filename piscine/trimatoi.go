package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	foundSign := false

	for _, r := range s {
		if r == '-' && !foundSign {
			sign = -1
			foundSign = true
		} else if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
			foundSign = true
		}
	}

	return result * sign
}
