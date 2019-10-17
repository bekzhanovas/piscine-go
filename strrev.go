package piscine

func StrRev(s string) string {
	last_string := ""
	for i := len (s) - 1; i>=0; i-- {
		last_string += string(s[i])
	}
	return last_string 
}
