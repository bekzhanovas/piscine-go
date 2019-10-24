package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	ans := make([]int, (max - min))
	for j := 0; j < (max - min); j++ {
		ans[j] = min + j
	}
	return ans
}
