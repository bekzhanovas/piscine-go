package piscine

func Capitalize(s string) string {
	s1 := []rune(s)
	Is_capital := 1
	for i := 0; i < StrLen2(s1); i++ {
		if is_capital(s1[i]) == 1 && Is_capital == 1 {
			if s1[i] >= 'a' && s1[i] <= 'z' {
				s1[i] = s1[i] - 32
			}
			Is_capital = 0

		} else if s1[i] >= 'A' && s1[i] <= 'Z' {
			s1[i] = s1[i] + 32
		} else if is_capital(s1[i]) == 0 {
			Is_capital = 1
		}
	}
	return string(s1)
}

func is_capital(a rune) int {
	if (a >= 'a' && a <= 'z') || (a >= 'A' && a <= 'Z') || (a >= '0' && a <= '9') {
		return 1
	}
	return 0
}
