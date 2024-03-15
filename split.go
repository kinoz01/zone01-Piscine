package main

import "fmt"

func Split(s, sep string) []string {
	startIndex := 0
	r := []string{}
	for i := 0; i+len(sep) <= len(s); i++ {
		if s[i:i+len(sep)] == sep { // if you find sep inside the string
			if s[startIndex:i] != "" { // don't append empty strings
				r = append(r, s[startIndex:i])
			}
			startIndex = i + len(sep)
			// i+=len(sep) // ADD THIS LINE TO INCLUDE THE "sep"
		} else if i+len(sep) == len(s) { // treating the case of the last word in case the string doesn't end with "sep"
			r = append(r, s[startIndex:])
		}
	}
	return r
}

func main() {
	s := "HAHelloHAHAHAHAhowHAHAareHAyou?HAHA"
	// s := "HelloHAhowHAareHAyou?" they apparently only tested this base case this is why "wrong" answer may also work
	fmt.Printf("%#v\n", Split(s, "HA"))
}

// quest 7 / checkpoint
