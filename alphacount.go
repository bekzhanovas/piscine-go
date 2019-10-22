package piscine

import "unicode"

func AlphaCount(s string) int {
	counter := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			counter++
		}
	}
	return counter
}
