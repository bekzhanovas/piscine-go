package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		last := n % 10
		n = n / 10
		z01.PrintRune(rune(last) + '0')
	}
}
