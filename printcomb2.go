package main

import "github.com/01-edu/z01"

func main() {
	PrintComb2()
}

func PrintComb2() {
	var o, o2, t, t2 int
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for x := '0'; x <= '9'; x++ {
				for y := '0'; y <= '9'; y++ {
					o = int(i) * 10
					t = int(j) * 1
					o2 = int(x) * 10
					t2 = int(y) * 1
					if t+o >= t2+o2 {
						continue
					}
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(x)
					z01.PrintRune(y)
					if i == '9' && j == '8' && x == '9' && y == '9' {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}

// o, o2 is for ones, t, t2 for tens
// ij number = i*10 + j*1
// xy number = x*10 + y*1
// quest 2
