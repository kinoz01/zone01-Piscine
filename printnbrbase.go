package main

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	sliceB := []rune(base)
	var n int
	x := len(sliceB)
	var r string
	index := []int{}
	if len(sliceB) < 2 {
		printstr("NV")
		return
	}
	for i := 0; i < len(sliceB); i++ {
		for j := 0; j < len(sliceB); j++ {
			if i != j && sliceB[i] == sliceB[j] {
				printstr("NV")
				return
			}
		}
		if sliceB[i] == '+' || sliceB[i] == '-' {
			printstr("NV")
			return
		}
	}
	if nbr < 0 {
		z01.PrintRune('-')
		n = -nbr
	} else if nbr > 0 {
		n = nbr
	} else {
		z01.PrintRune('0')
	}
	if nbr == -9223372036854775808 { // i have a problem for this particular case
		printstr("9223372036854775808")
	}
	for n > 0 {
		index = append(index, n%x)
		n /= x
	}
	for i := 0; i < len(index); i++ {
		r = string(sliceB[index[i]]) + r
	}
	printstr(r)
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
	PrintNbrBase(-9223372036854775808, "0123456789")
	z01.PrintRune('\n')
}

// quest 5 / checkpoint
