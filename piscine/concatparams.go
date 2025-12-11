package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// First calculate total size to allocate
	totalLen := 0
	for _, s := range args {
		totalLen += len(s)
	}
	totalLen += len(args) - 1 // for the '\n' between items

	// Allocate exactly once
	result := make([]byte, totalLen)

	// Fill the slice
	pos := 0
	for i, s := range args {
		// copy string bytes into result
		for j := 0; j < len(s); j++ {
			result[pos] = s[j]
			pos++
		}
		// add newline if it's NOT the last item
		if i != len(args)-1 {
			result[pos] = '\n'
			pos++
		}
	}

	return string(result)
}
