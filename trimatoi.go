package piscine

func TrimAtoi(s string) int {
	list := []rune(s)
	var int_list []int
	var int_list1 []int
	for i := 0; i < StrLen2(list); i++ {
		if list[i] >= '0' && list[i] <= '9' {
			int_list = append(int_list, int(list[i])-48)
		} else if list[i] == '+' || list[i] == '-' {
			int_list = append(int_list, int(list[i]))
		}
	}
	if StrLen1(int_list) > 1 && int_list[0] == 45 {
		int_list[1] = int_list[1] * (-1)
	}
	if StrLen1(int_list) > 1 {
		for j := 0; j < StrLen1(int_list); j++ {
			if int_list[j] != 43 && int_list[j] != 45 {
				int_list1 = append(int_list1, int_list[j])
			}
		}
	} else {
		return 0
	}
	final_int := 0
	if int_list1[0] >= 0 {
		for k := StrLen1(int_list1) - 1; k >= 0; k-- {
			final_int = final_int + int_list1[StrLen1(int_list1)-1-k]*RecursivePower1(10, k)
		}
	} else {
		for k := StrLen1(int_list1) - 1; k > 0; k-- {
			final_int = final_int + int_list1[StrLen1(int_list1)-k]*RecursivePower1(10, k-1)
		}
		final_int = int_list1[0]*RecursivePower1(10, StrLen1(int_list1)-1) - final_int
	}
	return final_int
}

func RecursivePower1(nb int, power int) int {
	if power == 1 {
		return nb
	} else if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	}
	return nb * RecursivePower1(nb, power-1)
}

func StrLen1(str1 []int) int {
	len_1 := 0
	for i := range str1 {
		len_1 = i + 1
	}
	return len_1
}

func StrLen2(str1 []rune) int {
	len_1 := 0
	for i := range str1 {
		len_1 = i + 1
	}
	return len_1
}
