package main

import "fmt"

func IsAlpha(s string) bool {
	sliceS := []rune(s)
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'A' && sliceS[i] <= 'Z' || sliceS[i] >= 'a' && sliceS[i] <= 'z' || sliceS[i] >= '0' && sliceS[i] <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
	fmt.Println(IsAlpha(""))
}

// quest 5
