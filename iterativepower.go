package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 1 {
		return nb
	}
	total := nb
	for i := 2; i <= power; i++ {
		total *= nb
	}
	return total
}
