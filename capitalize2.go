package main

import "fmt"

func Capitalize(s string) string {
	sliceS := []rune(s)
	capitalizeNext := true
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] <= 'z' && sliceS[i] >= 'a' && capitalizeNext {
			sliceS[i] -= 32
			capitalizeNext = false
		} else if sliceS[i] <= 'Z' && sliceS[i] >= 'A' && !capitalizeNext {
			sliceS[i] += 32
		}
		if sliceS[i] >= 'a' && sliceS[i] <= 'z' || (sliceS[i] >= 'A' && sliceS[i] <= 'Z') || (sliceS[i] <= '9' && sliceS[i] >= '0') {
			capitalizeNext = false
		} else {
			capitalizeNext = true
		}
	}
	return string(sliceS)
}

func main() {
	fmt.Println(Capitalize("HEAARello! AA HowAA are you? How+are+things+AA4AAyou?"))
}

// quest 5
