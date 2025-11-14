package piscine

func StrRev(s string) string {
	rev := []rune(s)
	n := len(rev)
	for a := 0; a < n/2; a++ {
		rev[a], rev[n-1-a] = rev[n-1-a], rev[a]
	}
	return string(rev)
}
