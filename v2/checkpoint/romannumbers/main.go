package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n := Atoi(os.Args[1])
	if n == 0 {
		os.Stdout.WriteString("ERROR: cannot convert to roman digit\n")
		return
	}
	romn := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roml := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var t []int
	var result string
	var formula string

	for i := 0; i < len(romn); i++ {
		t = append(t, n/romn[i])
		n %= romn[i]
	}
	for i := 0; i < len(roml); i++ {
		if t[i] == 0 {
			continue
		}
		for j := 0; j < t[i]; j++ {
			result += roml[i]
			if len(roml[i]) == 2 {
				formula += "(" + string(roml[i][1]) + "-" + string(roml[i][0]) + ")" + "+"
				continue
			}
			formula += roml[i] + "+"
		}
	}
	os.Stdout.WriteString(formula[:len(formula)-1] + "\n")
	os.Stdout.WriteString(result + "\n")
}

func Atoi(s string) int {
	var n int
	for _, char := range s {
		if char < '0' && char > '9' {
			return 0
		}
		n = n*10 + int(char-'0')
	}
	if n >= 4000 {
		return 0
	}
	return n
}
