package piscine

func IsPrintable(str string) bool {
	str1 := []rune(str)
	for i := 0; i < StrLen2(str1); i++ {
		if str1[i] >= 32 {
			continue
		} else {
			return false
		}
	}
	return true
}
