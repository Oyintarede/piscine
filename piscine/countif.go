package piscine

func CountIf(f func(s string) bool, a []string) int {
	count := 0
	for _, v := range a {
		if f(v) {
			count++
		}
	}
	return count
}
