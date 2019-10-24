package piscine

func SplitWhiteSpaces(str string) []string {
	var data []string
	str1 := []rune(str)
	min := 0
	max := 0
	counter := 0
	for _, letter := range str1 {
		max++
		counter++
		if letter == ' ' || letter == '\n' || letter == '\t' {
			data = append(data, str[min:max-1])
			min = max
		}
	}
	data = append(data, str[min:max])
	return data
}
