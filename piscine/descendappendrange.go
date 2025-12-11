package piscine

func DescendAppendRange(max, min int) []int {
	nums := []int{}

	for i := max; i > min; i-- {
		nums = append(nums, i)
	}

	return nums
}
