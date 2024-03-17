package main

import "github.com/01-edu/z01"

func Rot14(s string) string {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'a' && sliceS[i] <= 'l' || sliceS[i] >= 'A' && sliceS[i] <= 'L' {
			sliceS[i] += 14
		} else if sliceS[i] > 'l' && sliceS[i] <= 'z' || sliceS[i] > 'L' && sliceS[i] <= 'Z' {
			sliceS[i] -= 12
		}
	}
	return string(sliceS)
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

// hackathon
