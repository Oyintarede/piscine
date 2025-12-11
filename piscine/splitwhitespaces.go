package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	runes := []rune(s)

	start := -1 // start index of current word

	for i, r := range runes {
		// Check if current rune is whitespace
		if r == ' ' || r == '\t' || r == '\n' {
			if start != -1 {
				// End of a word
				words = append(words, string(runes[start:i]))
				start = -1
			}
		} else {
			// Start of a word
			if start == -1 {
				start = i
			}
		}
	}

	// End-of-string word
	if start != -1 {
		words = append(words, string(runes[start:]))
	}

	return words
}
