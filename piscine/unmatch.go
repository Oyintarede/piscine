package piscine

func Unmatch(a []int) int {
	var quant int
	for _, elems := range a {
		quant = 0
		for _, value := range a {
			if elems == value {
				quant++
			}
		}
		if quant%2 != 0 {
			return elems
		}
	}
	return -1
}
