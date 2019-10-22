package piscine

func Concat(str1 string, str2 string) string {
	str11 := []rune(str1)
	str21 := []rune(str2)
	str2_len := StrLen2(str21)
	for i := 0; i < str2_len; i++ {
		str11 = append(str11, str21[i])
	}
	return string(str11)
}
