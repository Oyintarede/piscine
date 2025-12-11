package piscine

func AtoiBase(s string, base string) int {
	baseLen := 0
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return 0
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
		baseLen++
	}
	if baseLen < 2 {
		return 0
	}

	result := 0

	for i := 0; i < len(s); i++ {
		idx := -1
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				idx = j
			}
		}
		if idx == -1 {
			return 0
		}
		result = result*baseLen + idx
	}

	return result
}
