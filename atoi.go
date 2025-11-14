package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sign := 1
	result := 0
	for i, r := range s {
		if i == 0 && (r == '+' || r == '-') {
			if r == '-' {
				sign = -1
			}
			continue
		}
		if r < '0' || r > '9' {
			return 0
		}

		result = result*10 + int(r-'0')
	}

	return result * sign
}
