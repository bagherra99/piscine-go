package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for i := 0; i <= len(s)-1; i++ {
		// fmt.Println(string(s[i]))
		z01.PrintRune(rune(s[i]))
		// z01.PrintRune(rune(10))
	}
}
