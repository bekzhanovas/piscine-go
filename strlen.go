package piscine

func StrLen(str string) int {
	len_1 := 0
	for i, _ := range str {
		len_1 = i + 1
	}
	return len_1
 }
