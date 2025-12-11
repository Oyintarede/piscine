package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}
	if start == len(s) {
		return 0
	}
	result := 0
	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0
		}
		result = result*10 + int(ch-'0')
	}
	return sign * result
}
