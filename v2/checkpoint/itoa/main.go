package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	s := ""
	var neg bool
	if n < 0 {
		neg = true
		n = -n
	}
	for n > 0 {
		s = string(rune(n%10)+'0') + s
		n /= 10
	}
	if neg {
		return "-" + s
	}
	return s
}
