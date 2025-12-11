package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	// First count how many numbers we need
	size := max - min

	// Create a slice using `var` and append (make is NOT allowed)
	var result []int

	// Append values from min to max-1
	for i := 0; i < size; i++ {
		result = append(result, min+i)
	}

	return result
}
