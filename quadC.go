package piscine

import "fmt"

//x represent le nombre de colonne et y le nombre de lignes
func QuadC(x int, y int) {
	//si x et y sont positifs la  fonction execute la boucle
	if x > 0 && y > 0 {
		if x > 1 && y > 1 {
			//cette boucle parcourt chaque ligne
			for i := 0; i < y; i++ {
				//cette boucle parcourt chaque colonne
				for j := 0; j < x; j++ {
					// ici le code
					if (i == 0 && j == 0) || (i == 0 && j == x-1) {
						fmt.Print("A")
					} else if (i == y-1 && j == x-1) || (i == y-1 && j == 0) {
						fmt.Print(`C`)
					} else if (i > 0 && i < y-1) && (j == 0 || j == x-1) {
						fmt.Print("B")
					} else if (j > 0 && j < x-1) && (i == 0 || i == y-1) {
						fmt.Print("B")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
		if x == 1 || y == 1 {
			//cette boucle parcourt chaque ligne
			for i := 0; i < y; i++ {
				//cette boucle parcourt chaque colonne
				for j := 0; j < x; j++ {
					// ici le code
					if i == 0 && j == 0 {
						fmt.Print("A")
					} else if i == 0 && j == x-1 {
						fmt.Print(`A`)
					} else if i == y-1 && j == 0 {
						fmt.Print(`C`)
					} else if (i > 0 && i < y-1) && (j == 0 || j == x-1) {
						fmt.Print("B")
					} else if (j > 0 && j < x-1) && (i == 0 || i == y-1) {
						fmt.Print("B")
					} else {
						fmt.Print(" ")
					}
				}
				fmt.Println()
			}
		}
	}
}
