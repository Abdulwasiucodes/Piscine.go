package piscine

import "github.com/01-edu/z01"

func z01print(r rune) {
	z01.PrintRune(r)
}

func DescendComb() {
	for z := '9'; z >= '0'; z-- {
		for y := '9'; y >= '0'; y-- {
			for x := '9'; x >= '0'; x-- {
				for w := '9'; w >= '0'; w-- {
					if (z > x) || (z == x && y > w) {
						z01print(z)
						z01print(y)
						z01print(' ')
						z01print(x)
						z01print(w)
						if z != '0' || y != '1' || x != '0' || w != '0' {
							z01print(',')
							z01print(' ')
						}
					}
				}
			}
		}
	}
}
