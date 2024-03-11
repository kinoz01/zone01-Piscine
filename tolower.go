package main

import "fmt"

func ToLower(s string) string {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'A' && sliceS[i] <= 'Z' {
			sliceS[i] += 32
		}
	}
	return string(sliceS)
}

func main() {
	fmt.Println(ToLower("Hello! How are you?"))
}

// quest 5
