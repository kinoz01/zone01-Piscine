package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	r := []string{}
	for _, byt := range arr {
		if byt == 0 {
			r = append(r, "00")
		} else {
			r = append(r, hex(int(byt)))
		}
	}

	// z01.PrintRune(rune(31)) 	// testing printable character min limits
	// z01.PrintRune(rune(127))	// testing printable character max limits

	// printing r:
	for i := 0; i < len(r); i++ {
		ppstr(r[i])
		z01.PrintRune(' ')
		if (i-3)%4 == 0 || i == len(r)-1 {
			z01.PrintRune('\n') // print new line each four strings and in the last one
		}
	}
	// printing ASCII graphic characters
	for _, byt := range arr {
		if byt < 32 || byt > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func ppstr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func hex(n int) string {
	base := "0123456789abcdef"
	sliceBase := []rune(base)
	index := []int{}
	var r []rune
	for n > 0 {
		index = append(index, n%16)
		n /= 16
	}
	for i := len(index) - 1; i >= 0; i-- {
		r = append(r, sliceBase[index[i]])
	}
	return string(r)
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// checkpoint / hardcoded using append (i don't know if it's allowed)
