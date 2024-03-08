package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

func PrintPositiveNum(n int) {
	// var r string
	c := '0'
	for i := 1; i <= n%10; i++ {
		c++
	}
	for i := -1; i >= n%10; i-- {
		c++
	}
	if n/10 != 0 {
		PrintPositiveNum(n / 10)
	}
	z01.PrintRune(c)
	/************if you want to use string fmt*******/
	// r=r+string(c)
	// fmt.Print(r)
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintPositiveNum(n)
}
