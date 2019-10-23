package main

import "os"
import "github.com/01-edu/z01"

func main() {
	argument := os.Args
	len_arg := StrLen2(argument)
	for i := 1; i < len_arg; i++ {
		a := []rune(argument[i])
		for _, j := range a {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
