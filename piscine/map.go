package piscine

func Map(f func(int) bool, a []int) []bool {
	var res []bool

	for _, v := range a {
		res = append(res, f(v))
	}

	return res
}
