package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var list []int
	if n == 0 {
		list = append(list, 0)
	}
	for n > 0 {
		last := n % 10
		n = n / 10
		list = append(list, last)
	}
	for {
		counter := 0
		for i := 1; i < len(list); i++ {
			if list[i] < list[i-1] {
				list[i], list[i-1] = list[i-1], list[i]
				counter++
			}
		}
		if counter == 0 {
			break
		}
	}
	for j := 0; j < len(list); j++ {
		z01.PrintRune(rune(list[j]) + '0')
	}
}
