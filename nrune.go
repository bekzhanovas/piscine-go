package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	if StrLen2(s1) < n || n < 1 {
		return '\x00'
	} else {
		return s1[n-1]
	}
}
