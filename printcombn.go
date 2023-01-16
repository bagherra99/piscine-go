package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	switch n {
	case 1:
		for i := 0; i <= 9; i++ {
			z01.PrintRune(rune(i) + 48)
		}
	case 2:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				z01.PrintRune(rune(i) + 48)
				z01.PrintRune(rune(j) + 48)
				if i != 8 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	case 3:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					z01.PrintRune(rune(i) + 48)
					z01.PrintRune(rune(j) + 48)
					z01.PrintRune(rune(k) + 48)
					if i != 7 {
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	case 4:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					for l := k + 1; l <= 9; l++ {
						z01.PrintRune(rune(i) + 48)
						z01.PrintRune(rune(j) + 48)
						z01.PrintRune(rune(k) + 48)
						z01.PrintRune(rune(l) + 48)
						if i != 6 {
							z01.PrintRune(44)
							z01.PrintRune(32)
						}
					}
				}
			}
		}
	case 5:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					for l := k + 1; l <= 9; l++ {
						for m := l + 1; m <= 9; m++ {
							z01.PrintRune(rune(i) + 48)
							z01.PrintRune(rune(j) + 48)
							z01.PrintRune(rune(k) + 48)
							z01.PrintRune(rune(l) + 48)
							z01.PrintRune(rune(m) + 48)
							if i != 5 {
								z01.PrintRune(44)
								z01.PrintRune(32)
							}
						}
					}
				}
			}
		}
	case 6:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					for l := k + 1; l <= 9; l++ {
						for m := l + 1; m <= 9; m++ {
							for n := m + 1; n <= 9; n++ {
								z01.PrintRune(rune(i) + 48)
								z01.PrintRune(rune(j) + 48)
								z01.PrintRune(rune(k) + 48)
								z01.PrintRune(rune(l) + 48)
								z01.PrintRune(rune(m) + 48)
								z01.PrintRune(rune(n) + 48)
								if i != 4 {
									z01.PrintRune(44)
									z01.PrintRune(32)
								}
							}
						}
					}
				}
			}
		}
	case 7:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					for l := k + 1; l <= 9; l++ {
						for m := l + 1; m <= 9; m++ {
							for n := m + 1; n <= 9; n++ {
								for o := n + 1; o <= 9; o++ {
									z01.PrintRune(rune(i) + 48)
									z01.PrintRune(rune(j) + 48)
									z01.PrintRune(rune(k) + 48)
									z01.PrintRune(rune(l) + 48)
									z01.PrintRune(rune(m) + 48)
									z01.PrintRune(rune(n) + 48)
									z01.PrintRune(rune(o) + 48)
									if i != 3 {
										z01.PrintRune(44)
										z01.PrintRune(32)
									}
								}
							}
						}
					}
				}
			}
		}
	case 8:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					for l := k + 1; l <= 9; l++ {
						for m := l + 1; m <= 9; m++ {
							for n := m + 1; n <= 9; n++ {
								for o := n + 1; o <= 9; o++ {
									for p := o + 1; p <= 9; p++ {
										z01.PrintRune(rune(i) + 48)
										z01.PrintRune(rune(j) + 48)
										z01.PrintRune(rune(k) + 48)
										z01.PrintRune(rune(l) + 48)
										z01.PrintRune(rune(m) + 48)
										z01.PrintRune(rune(n) + 48)
										z01.PrintRune(rune(o) + 48)
										z01.PrintRune(rune(p) + 48)
										if i != 2 {
											z01.PrintRune(44)
											z01.PrintRune(32)
										}
									}
								}
							}
						}
					}
				}
			}
		}
	case 9:
		for i := 0; i <= 9; i++ {
			for j := i + 1; j <= 9; j++ {
				for k := j + 1; k <= 9; k++ {
					for l := k + 1; l <= 9; l++ {
						for m := l + 1; m <= 9; m++ {
							for n := m + 1; n <= 9; n++ {
								for o := n + 1; o <= 9; o++ {
									for p := o + 1; p <= 9; p++ {
										for q := p + 1; q <= 9; q++ {
											z01.PrintRune(rune(i) + 48)
											z01.PrintRune(rune(j) + 48)
											z01.PrintRune(rune(k) + 48)
											z01.PrintRune(rune(l) + 48)
											z01.PrintRune(rune(m) + 48)
											z01.PrintRune(rune(n) + 48)
											z01.PrintRune(rune(o) + 48)
											z01.PrintRune(rune(p) + 48)
											z01.PrintRune(rune(q) + 48)
											if i != 1 {
												z01.PrintRune(44)
												z01.PrintRune(32)
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
