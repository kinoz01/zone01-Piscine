package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	PrintMemory([10]byte{' ', ' ', ' ', ' ', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	for i := range arr {
		// fmt.Println(byte(arr[i]))
		z01.PrintRune(rune(hex[arr[i]>>4]))
		z01.PrintRune(rune(hex[arr[i]&15]))

		if (i-3)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		} else {
			z01.PrintRune(' ')
		}

	}
	for _, bit := range arr {
		if bit < 33 || bit > 126 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(bit))
		}
	}
	z01.PrintRune('\n')
}
