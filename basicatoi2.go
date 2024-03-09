package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {
	var n int
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			return 0
		}
		n = n*10 + int(digit-'0')
	}
	return n
}

// quest 3
