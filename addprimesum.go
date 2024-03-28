package main

import "os"

func ateao(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

func iteao(n int) string {
	var r string
	if n == 0 {
		return "0"
	}
	for n > 0 {
		r = string(n%10+'0') + r
		n /= 10
	}
	return r
}

func prime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args[1:]) != 1 {
		os.Stdout.WriteString("0\n")
		return
	}
	s := os.Args[1]
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			os.Stdout.WriteString("0\n")
			return
		}
	}
	var r int
	for i := 0; i <= ateao(s); i++ {
		if prime(i) {
			r += i
		}
	}
	os.Stdout.WriteString(iteao(r) + "\n")
}

// checkpoint
