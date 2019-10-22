package piscine

func IsNumeric(str string) bool {
	str1 := []rune(str)
	for i := 0; i < StrLen2(str1); i++ {
		if str1[i] >= '0' && str1[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
