package main

import "os"
import "github.com/01-edu/z01"

func main() {
	argument := os.Args
	len_arg := 0
	for k := range argument {
		len_arg = k + 1
	}

	for i := len_arg - 1; i > 0; i-- {
		a := []rune(argument[i])
		for _, j := range a {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
