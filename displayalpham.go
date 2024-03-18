package main

import "github.com/01-edu/z01"

func main() {
	c := 0
	for i := 'a'; i <= 'z'; i++ {
		c++
		if c%2 == 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}

// checkpoint
