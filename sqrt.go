package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 0
	}
	if nb == 1 {
		return 1
	}
	for i := 2; i <= nb/2+1; i++ {
		if i*i == nb {
			return i
		}

		if i*i > nb {
			return 0
		}
	}

	return 0
}
