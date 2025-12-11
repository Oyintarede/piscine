package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""
	for _, char := range str + " " {
		if char == ' ' {
			if word != "" {
				summary[word]++
				word = ""
			} else {
				summary[""]++
			}
		} else {
			word += string(char)
		}
	}
	return summary
}
