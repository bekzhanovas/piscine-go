package main

import "github.com/01-ude/z01"

func main() {
	final := ""
	for a := 97; a < 123; a++ {
		character := rune(a)
		final += string(character)
	}
	fmt.Println(final)
}
