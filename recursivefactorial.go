package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 {
		return 0
	}
	return nb * IterativeFactorial(nb-1)
}
