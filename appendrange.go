package piscine

func AppendRange(min, max int) []int {
	var final []int
	if min >= max {
		return final
	}
	for i := min; i < max; i++ {
		final = append(final, i)
	}
	return final
}
