package piscine

func StrRev(s string) string {
	a := []rune(s)
	// return string(a)
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}
