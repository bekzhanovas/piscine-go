package piscine

func IsAlpha(str string) bool {
	str1 := []rune(str)
	for i := 0; i < StrLen2(str1); i++ {
		if (str1[i] >= '0' && str1[i] <= '9') || (str1[i] >= 'a' && str1[i] <= 'z') || (str1[i] >= 'A' && str1[i] <= 'Z') {
			continue
		} else {
			return false
		}
	}
	return true
}
