package main

import "github.com/01-edu/z01"

func main() {
	data := "./main"
	data1 := []rune(data)
	for i := 0; i < len(data1); i++ {
		z01.PrintRune(data1 [i])
	}
}
