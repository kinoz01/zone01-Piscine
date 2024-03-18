package main

import "fmt"

func Capitalize(s string) string {
	sliceS := []rune(s)
	capitalizeNext := true
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] <= 'Z' && sliceS[i] >= 'A' {
			sliceS[i] += 32 // make all lower case
		}
	}
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] <= 'z' && sliceS[i] >= 'a' && capitalizeNext {
			sliceS[i] -= 32
			capitalizeNext = false
		} else if (sliceS[i] >= 'a' && sliceS[i] <= 'z') || (sliceS[i] >= '1' && sliceS[i] <= '9') {
			capitalizeNext = false
		} else {
			capitalizeNext = true // if the current character is not 'A' to 'Z' nor 'a' to 'z' nor '0' to '9' capitalize the next
		}
	}
	return string(sliceS)
}

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

// quest 5 / checkpoint
