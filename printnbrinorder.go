package main

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var s string
	if n < 0 {
		return
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		s = string(n%10+'0') + s
		n /= 10
	}
	runeS := []rune(s)
	for i := 0; i < len(runeS)-1; i++ {
		for j := 0; j < len(runeS)-i-1; j++ {
			if runeS[j] > runeS[j+1] {
				runeS[j], runeS[j+1] = runeS[j+1], runeS[j]
			}
		}
	}
	for i := 0; i < len(runeS); i++ {
		z01.PrintRune(runeS[i])
	}
}

func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}

// quest 5
