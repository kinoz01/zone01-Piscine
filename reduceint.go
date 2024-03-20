package main

import "github.com/01-edu/z01"

func ReduceInt(a []int, f func(int, int) int) {
	result := f(a[0], a[1])
	printnum(result)
	z01.PrintRune('\n')
}

func printnum(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n/10 != 0 {
		printnum(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}

func main() {
	mul := func(acc int, cur int) int {
		return acc * cur
	}
	sum := func(acc int, cur int) int {
		return acc + cur
	}
	div := func(acc int, cur int) int {
		return acc / cur
	}
	as := []int{500, 2}
	ReduceInt(as, mul)
	ReduceInt(as, sum)
	ReduceInt(as, div)
}

// checkpoint
