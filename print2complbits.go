package main

import "os"

// import "fmt"

var exit bool

func ato(s string) int {
	var n int
	first := true
	for i, _ := range s {
		if (s[0] == '-' || s[0] == '+') && first {
			first = false
			continue
		}
		if s[i] > '9' || s[i] < '0' {
			exit = true
		}
		n = n*10 + int(s[i]-'0')
	}
	if s[0] == '-' {
		return -n
	}
	return n
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		os.Stdout.WriteString("Usage: <program> <number> <width>\n")
		return

	}
	n := ato(args[0])
	for _, digit := range args[1] {
		if digit > '9' || digit < '0' {
			os.Stdout.WriteString("Incorrect width\n")
			return
		}
	}
	bi := ato(args[1])
	var r string
	if exit || n == 0 {
		for i := 1; i <= bi; i++ {
			os.Stdout.WriteString("0")
			if i%4 == 0 {
				os.Stdout.WriteString(" ")
			}
		}
		os.Stdout.WriteString("\n")
		return
	}
	sign := false
	if n < 0 {
		n = -n
		sign = true
	}
	for n > 0 {
		r = string(n%2+'0') + r
		n /= 2
	}
	// fmt.Println(r)
	for i := len(r); i < bi; i++ {
		r = "0" + r
	}
	if sign {
		sliceR := []rune(r)
		for i := 0; i < len(sliceR); i++ {
			if sliceR[i] == '0' {
				sliceR[i] = '1'
			} else {
				sliceR[i] = '0'
			}
		}
		for i := len(sliceR) - 1; i >= 0; i-- {
			if sliceR[i] == '1' {
				sliceR[i] = '0'
			} else {
				sliceR[i] = '1'
				break
			}
		}
		r = string(sliceR)
		if sliceR[0] == '0' {
			r = "1" + r
		}
	}
	for i := 0; i < len(r); i++ {
		os.Stdout.WriteString(string(r[i]))
		if (i-3)%4 == 0 {
			os.Stdout.WriteString(" ")
		}
	}
	os.Stdout.WriteString("\n")
}

// edu

/**** usage
go run print2complbits.go 0 8
arg[0]-----> number
arg[1]-----> bit base

// when u give bit base out of range the program just ignore it
*****/

/********* Important
Notice that on limits like 128, 256 ... which is 1000 0000
we have also -128 which is also 1000 0000 this is why in
a 8 bits we can represent signed number only in the range
-128 to 127
********/

// https://www.allmath.com/twos-complement.php

/*Take a binary number 11111001 as an example. Without
TCR, we can interprate it as 2^7 + 2^6 + 2^5 + 2^4 +
2^3 + 1 = 249. With TCR (i.e. applying the mathematical
operation), we view it as -2^7 + 2^6 + 2^5 + 2^4 + 2^3 +
1 = -7.

Similiary, for another binary number, say 00000111.
Without TCR, we view it as 2^2 + 2^1 + 1 = 7. And with
TCR, we view it as 2^2 + 2^1 + 1 = 7, which is identical
to the former one.
*/
