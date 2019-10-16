package main

import "github.com/01-edu/z01"

func main() {
	final := ""
	for a := 97; a < 123; a++ {
		character := rune(a)
		final += string(character)
	}
	z01.Println(final)
}
