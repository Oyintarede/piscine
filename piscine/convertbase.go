package piscine

func indexInBase(r rune, base []rune) int {
	for i, b := range base {
		if b == r {
			return i
		}
	}
	return -1
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	from := []rune(baseFrom)
	to := []rune(baseTo)
	nbrRunes := []rune(nbr)

	baseF := len(from)
	baseT := len(to)

	// -------- STEP 1: Convert nbr → base 10 integer --------
	value := 0
	for _, r := range nbrRunes {
		d := indexInBase(r, from)
		value = value*baseF + d
	}

	// -------- STEP 2: Convert integer → baseTo --------
	if value == 0 {
		return string(to[0])
	}

	var digits []rune

	for value > 0 {
		digits = append(digits, to[value%baseT])
		value /= baseT
	}

	// reverse digits
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return string(digits)
}
