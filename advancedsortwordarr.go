package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	if len(a) < 2 {
		return
	}
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
