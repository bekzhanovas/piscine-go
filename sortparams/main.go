package main

import "os"
import "github.com/01-edu/z01"

func main() {
	argument := os.Args[1:]
	len_arg := 0
	for k := range argument {
		len_arg = k + 1
	}
	for {
		counter := 0
		for i := 0; i < len_arg; i++ {
			if argument[i] > argument[i+1] {
				argument[i], argument[i+1] = argument[i+1], argument[i]
				counter++
			}
		}
		if counter == 0 {
			break
		}
	}
	for i := 0; i < len_arg; i++ {
		a := []rune(argument[i])
		for _, j := range a {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
