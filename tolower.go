package piscine

func ToLower(s string) string {
	s1 := []rune(s)
	len_s := len(s1)
	for i := 0; i < len_s; i++ {
		if s1[i] >= 'A' && s1[i] <= 'Z' {
			s1[i] = s1[i] + 32
		}

	}
	return string(s1)
}
