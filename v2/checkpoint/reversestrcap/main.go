package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		argRune := []rune(arg)
		for i, char := range argRune {
			if i+1 == len(argRune) || argRune[i+1] == ' ' {
				if char >= 'a' && char <= 'z' {
					argRune[i] -= 32
				}
			} else if char >= 'A' && char <= 'Z' {
				argRune[i] += 32
			}
		}
		os.Stdout.WriteString(string(argRune) + "\n")
	}
}
