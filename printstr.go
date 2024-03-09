package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _ , char := range s{
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

func main() {
	PrintStr("Hello World!")
}

// quest 3 
