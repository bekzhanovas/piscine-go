package piscine

func SplitWhiteSpaces(str string) []string {
	str1 := []rune(str)
	min := 0
	max := 0
	counter := 1
	for _, i := range str1 {
		if i == ' ' || i == '\n' || i == '\t' {
			counter++
		}
	}
	data := make([]string, counter)
	letter_number := 0
	for _, letter := range str1 {
		max++
		if letter == ' ' || letter == '\n' || letter == '\t' {
			data[letter_number] = str[min : max-1]
			min = max
			letter_number++
		}
	}
	data[letter_number] = str[min:max]
	return data
}
