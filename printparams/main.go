package main

import "os"
import "github.com/01-edu/z01"

func main() {
	argument := []string{"askar", "bar", "dat"}
	len_arg := 0
	for k := range argument {
		len_arg = k + 1
	}
	for i := 1; i < len_arg; i++ {
		a := []rune(argument[i])
		for _, j := range a {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
