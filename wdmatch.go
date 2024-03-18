package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	s1 := []rune(args[0])
	s2 := []rune(args[1])
	if len(s1) > len(s2) {
		return
	}
	c := 0
	for i := 0; i < len(s1); i++ {
		j := c // we gonna start our loop (searching) from where we found the last character (c=j+1)
		found := false
		for j < len(s2) {
			if s1[i] == s2[j] {
				c = j + 1
				found = true
				break
			} else {
				j++
			}
		}
		if !found { // if just one character is not found that's it return
			return
		}
	}
	os.Stdout.WriteString(args[0] + "\n")
}

// checkpoint
