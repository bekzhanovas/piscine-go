package piscine

func NRune(s string, n int) rune {
	s1 := []rune(s)
	if StrLen2(s1) < n {
		return '0'
	} else {
		return s1[n-1]
	}
}
