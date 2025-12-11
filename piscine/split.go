package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	runes := []rune(s)
	sepr := []rune(sep)

	var result []string

	n := len(runes)
	m := len(sepr)

	start := 0

	for i := 0; i <= n-m; {
		match := true
		for j := 0; j < m; j++ {
			if runes[i+j] != sepr[j] {
				match = false
				break
			}
		}

		if match {
			// Add substring before separator
			result = append(result, string(runes[start:i]))
			i += m    // skip the separator
			start = i // move start
		} else {
			i++
		}
	}

	// Append the final chunk
	result = append(result, string(runes[start:]))

	return result
}
