package main

import "github.com/01-edu/z01"

func main() {
	for i := '9'; i >= '0'; i-- {
		for j := '9'; j >= '0'; j-- {
			for k := '9'; k >= '0'; k-- {
				if i <= j || j <= k {
					continue
				}
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i == '2' && j == '1' && k == '0' {
					break
				}
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
