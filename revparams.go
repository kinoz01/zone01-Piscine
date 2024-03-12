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
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		printstr(args[i])
	}
}

// quest 6
