package piscine

import "fmt"

//x represent le nombre de colonne et y le nombre de lignes
func QuadA(x int, y int) {
	//si x et y sont positifs la  fonction execute la boucle
	if x > 0 && y > 0 {
		//cette boucle parcourt chaque ligne 
		for i := 0; i < y; i++ {
			//cette boucle parcourt chaque colonne
			for j := 0; j < x; j++ {
				// ici le code
				if (i == 0 || i == y-1) && (j == 0 || j == x-1) {
					fmt.Print("o")
				} else if (i > 0 && i < y-1) && (j == 0 || j == x-1) {
					fmt.Print("|")
				} else if (j > 0 && j < x-1) && (i == 0 || i == y-1) {
					fmt.Print("-")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	}
}