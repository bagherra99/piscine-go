package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := 0; j <= 8; j++ {
			for k := 0; k <= 9; k++ {
				z01.PrintRune(rune(i) + 48)
				z01.PrintRune(rune(j) + 48)
				z01.PrintRune(rune(k) + 48)
				if !(i == 7 && j == 8 && k == 9) {
					z01.PrintRune(rune(32))
					z01.PrintRune(rune(44))
				}
			}
		}
	}
}
