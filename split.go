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
		} else if i+len(sep) == len(s) { // if len s[i:i+len(sp)]!= sep and i+len(sp)=len(s) (WHIICH MEANS we are at the last iteration) WHIICH MEANS if len s[len(s)-len(sep):len(s))]!= sep WHIICH MEANS if the end of the string doesn't end with a sep
			r = append(r, s[startIndex:]) // in this case append all to the end
		}
	}
	return r
}

func main() {
	s := "HAHelloHAHAHAHAhowHAHAareHAyou?HAHA"
	// s := "HelloHAhowHAareHAyou?" they apparently only tested this base case this is why "wrong" answer may also work
	fmt.Printf("%#v\n", Split(s, "HA")) // this line allow us to see clearly how and why we added the condition "if s[index:i]!=""{" above
}

// quest 7 / checkpoint
