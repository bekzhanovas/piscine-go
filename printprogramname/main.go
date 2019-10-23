package piscine

import "os"
import "github.com/01-edu/z01"

func main() {
	data := os.Args[0]
	for i := range data {
		z01.PrintRune(i)
	}
}
