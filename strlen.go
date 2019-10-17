package piscine

func StrLen(str string) int {
	len_1 := 0
	str1 := []rune(str)
	for i := range str1 {
		len_1 = i + 1
	}
	return len_1
}
