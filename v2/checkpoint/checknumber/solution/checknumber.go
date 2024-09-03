package solution

func CheckNumber(arg string) bool {
	for _, c := range arg {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
}