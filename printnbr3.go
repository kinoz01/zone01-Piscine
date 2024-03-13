package main

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	printnbr0(n)
}

// we had to call a new function for recursion just to handle the case of 0
func printnbr0(n int) {
	nbr := n

	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -n
	}
	if nbr > 0 {
		d := nbr % 10
		printnbr0(nbr / 10)
		z01.PrintRune(rune(d+48)) // d int to rune conversion
	}
}

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
	// z01.PrintRune(rune(78))
	// z01.PrintRune(rune(48))
	// z01.PrintRune(rune(39))
}

// Quest 2 (using recursion)
