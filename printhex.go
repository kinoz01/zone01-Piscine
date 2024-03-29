package main

import "os"

func atoip(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	s := os.Args[1]
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			os.Stdout.WriteString("ERROR\n")
			return
		}
	}
	n := atoip(s)
	hex := "0123456789abcdef"
	ind := []int{}
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	for n > 0 {
		ind = append(ind, n%16)
		n /= 16
	}
	for i := len(ind) - 1; i >= 0; i-- {
		os.Stdout.WriteString(string(hex[ind[i]]))
	}
	os.Stdout.WriteString("\n")
}

// checkpoint
