package piscine

func TrimAtoi(s string) []int {
	list := []rune(s)
	var int_list []int
	var int_list1 []int
	for i := 0; i < len(list); i++ {
		if list[i] >= '0' && list[i] <= '9' {
			int_list = append(int_list, int(list[i])-48)
		} else if list[i] == '+' || list[i] == '-' {
			int_list = append(int_list, int(list[i]))
		}
	}
	if len(int_list) > 1 && int_list[0] == 45 {
		int_list[1] = int_list[1] * (-1)
	}
	if len(int_list) > 1 {
		for j := 0; j < len(int_list); j++ {
			if int_list[j] != 43 && int_list[j] != 45 {
				int_list1 = append(int_list1, int_list[j])
			}
		}
	}
	final_int := 0
	if int_list1[0] > 0 {
		for k := len(int_list1) - 1; k >= 0; k-- {
			final_int = final_int + int_list1[len(int_list1)-1-k]*RecursivePower1(10, k)
		}
	} else {
		for k := len(int_list1) - 1; k > 0; k-- {
			final_int = final_int + int_list1[len(int_list1)-k]*RecursivePower1(10, k-1)
		}
		final_int = int_list1[0]*RecursivePower1(10, len(int_list1)-1) - final_int
	}
	fmt.Println(final_int)
	return int_list1
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
