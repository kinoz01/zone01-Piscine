package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	s := []rune(args[0])
	if isvowel(s[0]) {
		os.Stdout.WriteString(string(s) + "ay\n")
		return
	}
	for i := 1; i < len(s); i++ {
		if isvowel(s[i]) {
			os.Stdout.WriteString(string(s[i:]) + string(s[0:i]) + "ay\n")
			return
		}
	}
	os.Stdout.WriteString("No vowels\n")	
}

func isvowel(r rune) bool {
	v := "aeiouAEIOU"
	for i := 0; i < len(v); i++ {
		if r == rune(v[i]) {
			return true
		}
	}
	return false
}

// checkpoint
