package piscine

func SortWordArr(a []string) {
	if len(a) < 2 {
		return
	}
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
