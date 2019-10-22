package piscine

func IsUpper(str string) bool {
	str1 := []rune(str)
	for i := 0; i < StrLen2(str1); i++ {
		if str1[i] >= 'A' && str1[i] <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
