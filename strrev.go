package piscine

func dfgdfg(str string) int {
	len_1 := 0
	str1 := []rune(str)
	for i := range str1 {
		len_1 = i + 1
	}
	return len_1
}

func StrRev(s string) string {
	last_string := ""
	len_s := dfgdfg(s)
	for i := len_s - 1; i >= 0; i-- {
		last_string += string(s[i])
	}
	return last_string
}
