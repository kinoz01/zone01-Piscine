package main

import "fmt"

func ToUpper(s string) string {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'a' && sliceS[i] <= 'z' {
			sliceS[i] -= 32
		}
	}
	return string(sliceS)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}

// quest 5
