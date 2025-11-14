package piscine

func Max(a []int) int {
	max := a[0]
	for y := range a {
		if max < a[y] {
			max = a[y]
		}
	}
	return max
}
