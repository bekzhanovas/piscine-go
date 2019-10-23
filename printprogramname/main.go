package piscine

import "os"
import "github.com/01-edu/z01"

func main() {
	a := os.Args[0]
	p := []rune(a)
	for _, l := range p {
		z01.PrintRune(l)
	}
	z01.PrintRune('\n')
}
