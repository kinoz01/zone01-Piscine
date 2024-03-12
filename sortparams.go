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
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for _, param := range args {
		printstr(param)
	}
}

// quest 6
