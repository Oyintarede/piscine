package piscine

func Compare(a, b string) int {
	aRunes := []rune(a)
	bRunes := []rune(b)

	minLen := len(aRunes)
	if len(bRunes) < minLen {
		minLen = len(bRunes)
	}

	for i := 0; i < minLen; i++ {
		if aRunes[i] < bRunes[i] {
			return -1
		}
		if aRunes[i] > bRunes[i] {
			return 1
		}
	}

	if len(aRunes) < len(bRunes) {
		return -1
	}
	if len(aRunes) > len(bRunes) {
		return 1
	}

	return 0
}
