package main

import "os"

func main() {
	if len(os.Args[1:]) != 1 {
		os.Stdout.WriteString("\n")
		return
	}
	s := os.Args[1]
	r := []rune{}
	c := false
	for _, char := range s {
		if char != ' ' {
			r = append(r, char)
			c = true
		} else if c {
			r = append(r, char)
			c = false
		}
	}
	os.Stdout.WriteString(string(r) + "\n")
}

// checkpoint
