package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printstr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func main() {
	// name := os.Args[0] ; n = name[2:] ; printstr(n)
	name := os.Args[0][2:] // [2:] remove ./ from program name
	printstr(name)
}

// quest 6
