package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	r := []string{}
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if index != i {
				r = append(r, s[index:i])
			}
			index = i + 1
		} else if i == len(s)-1 { // else if s[i]!=' ' && we are at the last character
			r = append(r, s[index:]) // append everything until the end of the string
		}
	}
	return r
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

// quest 7
