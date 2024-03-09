package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(NRune("Hello!", 3))
	z01.PrintRune(NRune("Salut!", 2))
	z01.PrintRune(NRune("Bye!", -1))
	z01.PrintRune(NRune("Bye!", 5))
	z01.PrintRune(NRune("Ola!", 4))
	z01.PrintRune('\n')
}

func NRune(s string, n int) rune {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if i == n-1 {
			return sliceS[i]
		}
	}
	return 0
}

/******************** another way *********************/
// func NRune(s string, n int) rune {
// 	sliceS := []rune(s)
// 	if n <= 0 {
// 		return 0
// 	} else if n > len(sliceS) {
// 		return 0
// 	} else if n == len(sliceS) {
// 		return sliceS[len(sliceS)-1]
// 	} else {
// 		return sliceS[n-1]
// 	}
// }

// quest 5
