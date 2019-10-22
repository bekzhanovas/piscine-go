package piscine

func Index(s string, toFind string) int {
	s1 := []rune(s)
	toFind1 := []rune(toFind)
	for i := 0; i <= (StrLen2(s1) - StrLen2(toFind1)); i++ {
		if toFind == s[i:i+StrLen2(toFind1)] {
			return i
		}
	}
	return -1
}
