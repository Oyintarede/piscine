package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}

	sLen := len(s)
	fLen := len(toFind)

	if fLen > sLen {
		return -1
	}

	for i := 0; i <= sLen-fLen; i++ {
		match := true
		for j := 0; j < fLen; j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
