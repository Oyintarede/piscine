package piscine

func f(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) <= 1 {
		return true
	}

	asc := true
	desc := true

	for i := 1; i < len(a); i++ {
		comp := f(a[i-1], a[i])
		if comp > 0 {
			asc = false
		}
		if comp < 0 {
			desc = false
		}
	}

	return asc || desc
}
