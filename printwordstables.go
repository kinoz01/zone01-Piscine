package main

import "github.com/01-edu/z01"

func SplitWhiteSpaces(s string) []string {
	r := []string{}
	index := 0 // representing start of words
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if s[index] != ' ' { // Don't append white spaces
				r = append(r, s[index:i])
			}
			index = i + 1
		} else if i == len(s)-1 { // else if s[i]!=' ' && we are at the last character
			r = append(r, s[index:]) // append everything until the end of the string
		}
	}
	return r
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func PrintWordsTables(a []string) {
	for _, str := range a {
		printstr(str)
	}
}

func main() {
	a := SplitWhiteSpaces("Hello how are you?")
	PrintWordsTables(a)
}

// quest 7
