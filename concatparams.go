package piscine

func ConcatParams(args []string) string {
	dl := 0
	for i := range args {
		dl = i + 1
	}
	ans := make([]string, dl)
	for j := 0; j < (dl - 1); j++ {
		ans[1] = ans[1] + args[j] + "\n"
	}
	ans[1] = ans[1] + args[(dl-1)]
	return ans[1]
}
