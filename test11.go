package main

import (
	"fmt"
	//"strconv"
	//"cast/ToInt"
)

func main() {
	str1 := ".96.4...1"
	str2 := "1...6...4"
	str3 := "5.481.39."
	str4 := "..795..43"
	str5 := ".3..8...."
	str6 := "4.5.23.18"
	str7 := ".1.63..59"
	str8 := ".59.7.83."
	str9 := "..359...7"
	raw_data := [9]string{str1, str2, str3, str4, str5, str6, str7, str8, str9}
	//переводим точки на 0 и в int
	for i, numbers := range raw_data {
		raw_data[i] = dots_to_zero(numbers)
	}
	////переводим string в int
	int_data := str_to_int(raw_data)

	//fmt.Println(int_data)
	//fmt.Println(int_data[1][0])
	
	for row :=range int_data {
		fmt.Println(int_data [row])
	}
}

func dots_to_zero(str string) string {
	str1 := []rune(str)
	for i, letter := range str1 {
		if letter == '.' {
			str1[i] = '0'
		}
	}
	return string(str1)
}

func str_to_int(s [9]string) [][]int {
	//s := []string{"096040001", "096040001"}
	int_list_all := make([][]int, 9)

	for number, str := range s {
		int_list := make([]int, 9)
		list := []rune(str)
		for i := 0; i < 9; i++ {
			int_list[i] = int(list[i]) - 48
		}
		int_list_all[number] = int_list
	}

	return int_list_all
}




