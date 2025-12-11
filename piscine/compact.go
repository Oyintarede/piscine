package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	newSlice := []string{}
	for _, v := range slice {
		if v != "" {
			newSlice = append(newSlice, v)
		}
	}
	*ptr = newSlice
	return len(newSlice)
}
