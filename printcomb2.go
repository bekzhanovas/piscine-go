package piscine

import "github.com/01-edu/z01"

func get_l(i, j, k rune) (rune) {
	if k == i {
		return j+1
	} else {
		return '0'
	}
}


func PrintComb2() {
	
	for i := '0'; i <= '9'; i++ {
		
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				l := get_l (i, j, k)
				for ; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(32)
					z01.PrintRune(k)
					z01.PrintRune(l)
					if i != '9' || j != '8' || k != '9' || l != '9' {
						z01.PrintRune(44)
						z01.PrintRune(32)
					} else {
						z01.PrintRune(10)
					}
	
				}
			}
		}
	}
}