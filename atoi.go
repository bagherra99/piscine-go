package piscine

func Atoi(s string) int {
	var result int
	var foundDigit bool
	var sign int = 1 // Default sign is positive
	var foundSign bool
	for i, c := range s {
		if c == '-' && !foundDigit && !foundSign {
			sign = -1
			foundSign = true
		} else if c == '+' && !foundDigit && !foundSign {
			sign = 1
			foundSign = true
		} else if c >= '0' && c <= '9' {
			result = result*10 + int(c-'0')
			foundDigit = true
		} else if foundDigit {
			for _, digit := range s[i:] {
				if digit != ' ' {
					return 0
				}
			}
			return result * sign
		} else if c != ' ' && !foundDigit {
			return 0
		}
	}
	return result * sign
}
