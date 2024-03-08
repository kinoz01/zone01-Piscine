package main

import "github.com/01-edu/z01"

func main() {
	PrintComb()
}

func PrintComb() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i == j && j == k {
					continue
				}
				if i >= j || j >= k {
					continue
				}
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)
				if i != '7' || j != '8' || k != '9' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// method using i!='7' || j!='8' || k!='9' to exclude 789
//
// quest 2.
