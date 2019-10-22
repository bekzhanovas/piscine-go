package piscine

func LastRune(s string) rune {
	s1 := []rune(s)
	return s1[StrLen2(s1)-1]
}
