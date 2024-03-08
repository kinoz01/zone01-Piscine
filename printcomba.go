package main

import "github.com/01-edu/z01"

func main() {
	PrintComba()
}

func PrintComba() {
	for i := '0'; i <= '7'; i++ {
		for j := i + 1; j <= '8'; j++ {
			for k := j + 1; k <= '9'; k++ {
				if k == '9' && i == '7' && j == '8' {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
				} else {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(k)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// using the if i>=j and j>=k while initializing loop incrementals
