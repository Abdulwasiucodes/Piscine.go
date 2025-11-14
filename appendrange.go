package piscine

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}

	var result []int
	for x := min; x < max; x++ {
		result = append(result, x)
	}
	return result
}
