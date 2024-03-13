package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoii(s string) int {
	sliceS := []rune(s)
	var n int
	for i := 0; i < len(sliceS); i++ {
		n = n*10 + int(sliceS[i]-'0')
	}
	return n
}

func main() {
	args := os.Args[1:]
	first := true
	for i := 0; i < len(args); i++ {
		if args[0] == "--upper" && first {
			first = false
			continue
		} else if atoii(args[i]) > 26 || atoii(args[i]) < 1 {
			z01.PrintRune(' ')
			continue
		} else if !first {
			z01.PrintRune(rune(atoii(args[i]) + 64))
		} else {
			z01.PrintRune(rune(atoii(args[i]) + 96))
		}
	}
	z01.PrintRune('\n')
}

// quest 6
