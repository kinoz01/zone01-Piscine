package main

import "github.com/01-edu/z01"

func main() {
	pstr("Hello World!\n")
}

func pstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

// checkpoint
