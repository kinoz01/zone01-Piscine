package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	sliceS := []rune(s)
	return sliceS[len(sliceS)-1]
}

func main() {
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Ŝalut!ẑ"))
	z01.PrintRune(LastRune("Ola!a"))
	z01.PrintRune('\n')
}

// quest 5
