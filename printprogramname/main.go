package main

import "os"
import "github.com/01-edu/z01"

func main() {
	argument := os.Args[0]
	a := []rune(argument)
	for _, i := range a {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
