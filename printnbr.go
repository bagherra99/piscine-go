package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	// if n < 0 {
	// 	z01.PrintRune('-')
	// 	n = -n
	// }

	// if n == 0 {
	// 	z01.PrintRune('0')
	// }
	j := 1
	if n < 0 {
		z01.PrintRune('-')
		j = -1
	}

	if n != 0 {
		quotient := (n / 10) * j
		if quotient != 0 {
			PrintNbr(quotient)
		}
	}
	reste := (n % 10) * j
	z01.PrintRune(rune(reste) + 48)
}
