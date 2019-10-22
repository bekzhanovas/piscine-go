package piscine

func IsLower(str string) bool {
	str1 := []rune(str)
	for i := 0; i < StrLen2(str1); i++ {
		if str1[i] >= 'a' && str1[i] <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
