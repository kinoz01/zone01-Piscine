package main

import (
	"fmt"
	"strings"
)

func Split(s, sep string) []string {
	startIndex := 0
	r := []string{}

	// Handle empty separator
	if sep == "" {
		for _, ch := range s {
			r = append(r, string(ch))
		}
		return r
	}

	for i := 0; i+len(sep) <= len(s); {
		if s[i:i+len(sep)] == sep { // if you find sep inside the string
			r = append(r, s[startIndex:i])
			startIndex = i + len(sep)
			i = startIndex // Move i forward to avoid overlapping separators
		} else {
			i++
		}
	}

	// Append the remaining substring after the loop
	r = append(r, s[startIndex:])
	return r
}

func main() {
	s := "HAHelloHAHAHAHAhowHAHAareHAyou?HAHA"
	// s := "HelloHAhowHAareHAyou?" they apparently only tested this base case this is why "wrong" answer may also work

	fmt.Printf("%#v\n", Split(s, "HA")) // this line allow us to see clearly how and why we added the condition "if s[index:i]!=""{" above
	fmt.Println(Split("a b c", " "))
	fmt.Println(Split("ggg - ddd - b", " - "))
	fmt.Println(Split("ee,ff,g,", ","))
	fmt.Println(Split("Riad", " "))
	fmt.Println(Split("rrrr", "rr"))
	fmt.Println(Split("rrirr", "rr"))
	fmt.Println(Split("Riad", ""))
	fmt.Println(Split("l", "ll"))
	fmt.Println(len(strings.Split("rrrr", "rr")))
}

// quest 7 / checkpoint
