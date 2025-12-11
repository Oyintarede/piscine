package piscine

func Abort(a, b, c, d, e int) int {
	values := []int{a, b, c, d, e}
	values = sortList(values)
	return values[2]
}

func sortList(params []int) []int {
	for i := 0; i < len(params)-1; i++ {
		for j := 0; j < len(params)-1-i; j++ {
			if params[j] > params[j+1] {
				params[j], params[j+1] = params[j+1], params[j]
			}
		}
	}
	return params
}
