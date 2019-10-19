package piscine


func IterativeFactorial(nb int) int {
	if nb > 0 {
		Total := 1
		for i := 1; i <= nb; i++ {
			Total *= i
		}
		return Total
	} else {
		return 0
	}
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}
