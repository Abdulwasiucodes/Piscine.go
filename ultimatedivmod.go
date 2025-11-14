package piscine

func UltimateDivMod(a *int, b *int) {
	div := *a / *b
	rem := *a % *b
	*a = div
	*b = rem
}
