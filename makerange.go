package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min
	result := make([]int, size)

	for x := 0; x < size; x++ {
		result[x] = min + x
	}

	return result
}
