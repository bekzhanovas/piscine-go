package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	return s1[n-1]
}
