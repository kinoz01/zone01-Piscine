package main

import "github.com/01-edu/z01"

func FirstRune(s string) rune {
	sliceS := []rune(s)
	return sliceS[0]
}

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Ŝalut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

// quest 5
