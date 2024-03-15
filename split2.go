package main

import "fmt"

// // adapting the previous splitwhitespaces to the new split format
func SplitWhiteSpaces2(s string) []string {
	r := []string{}
	startIndex := 0
	for i := 0; i+1 <= len(s); i++ {
		if s[i:i+1] == " " {
			if s[startIndex:i] != "" {
				r = append(r, s[startIndex:i])
			}
			startIndex = i + 1
			// i+=1 // ADD THIS LINE TO INCLUDE THE "sep"
		} else if i+1 == len(s) {
			r = append(r, s[startIndex:])
		}
	}
	return r
}

// split with append usage restriction
func Split2(s, sep string) []string {
	var result []string
	count := 0
	start := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			count++
		}
	}
	count++
	result = make([]string, count)
	index := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result[index] = s[start:i]
			start = i + len(sep)
			index++
			i += len(sep) // ADD THIS LINE TO INCLUDE THE "sep"
		} else if i+len(sep) == len(s) {
			result[index] = s[start:]
		}
	}
	// removing empty strings :
	c2 := 0
	indexx := 0
	for i := 0; i < len(result); i++ {
		if result[i] != "" {
			c2++
		}
	}
	final := make([]string, c2)
	for i := 0; i < len(result); i++ {
		if result[i] == "" {
			continue
		}
		final[indexx] = result[i]
		indexx++
	}
	return final
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces2("Hello how are you?"))
	s := "HAHelloHAHAHAhowHAHAareHAyou?HAHAHA"
	//d := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split2(s, "HA"))
}

// quest 7 / checkpoint
