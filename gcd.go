package main

import "os"

var exit bool

func atoip(s string) int {
	var n int
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			exit = true
		}
	}
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	if n == 0 {
		exit = true
	}
	return n
}

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}
	n1 := atoip(os.Args[1])
	n2 := atoip(os.Args[2])
	var gcd int
	if exit {
		return
	}
	for i := 1; i <= n2; i++ {
		if n1%i == 0 && n2%i == 0 {
			gcd = i
		}
	}
	var r string
	for gcd > 0 {
		r = string(gcd%10+'0') + r
		gcd /= 10
	}
	os.Stdout.WriteString(r + "\n")
}

// checkpoint
