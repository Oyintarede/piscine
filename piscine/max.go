package piscine

func Max(a []int) int {
	sortA := sortListNum(a)

	return sortA[len(sortA)-1]
}

func sortListNum(params []int) []int {
	for i := 0; i < len(params)-1; i++ {
		for j := 0; j < len(params)-1-i; j++ {
			if params[j] > params[j+1] {
				params[j], params[j+1] = params[j+1], params[j]
			}
		}
	}
	return params
}
