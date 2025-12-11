package piscine

func StringToIntSlice(str string) []int {
	runes := []rune(str)

	result := []int(nil)
	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}
	return result
}
