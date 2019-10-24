package piscine

func ConcatParams(args []string) string {
	str := ""
	for _, i := range args {
		str = str + "`" + i + "\n"
	}
	return str
}
