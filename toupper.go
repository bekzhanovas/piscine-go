package piscine

func ToUpper(s string) string {
	s1 := []rune(s)
	len_s := len(s1)
	for i := 0; i < len_s; i++ {
		if s1[i] >= 'a' && s1[i] <= 'z' {
			s1[i] = s1[i] - 32
		}

	}
	return string(s1)
}
