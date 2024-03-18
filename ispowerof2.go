package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	s := args[0]
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			return
		}
	}
	/************ Overflow Check **************/
	if len(s) > 19 || len(s) == 19 && s[0] >= '9' && s[1] >= '2' && s[2] >= '2' && s[3] >= '3' && s[4] >= '3' && s[5] >= '7' && s[6] >= '2' && s[7] >= '0' && s[8] >= '3' && s[9] >= '6' && s[10] >= '8' && s[11] >= '5' && s[12] >= '4' && s[13] >= '7' && s[14] >= '7' && s[15] >= '5' && s[16] >= '8' && s[17] >= '0' && s[18] > '7' {
		return
	}
	///
	n := atoio(s)
	// checking if n can be represented as n = 2^(m)
	if n == 0 {
		os.Stdout.WriteString("false\n")
		return
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
			continue
		} else {
			os.Stdout.WriteString("false\n")
			return
		}
	}
	os.Stdout.WriteString("true\n")

	/********* Another Way for Checking**********/
	// p := 2
	// if n == 1 {
	// 	os.Stdout.WriteString("true\n") // 1 = 2⁰
	// 	return
	// }
	// for p < n {
	// 	p = p * 2
	// }
	// if p == n {
	// 	os.Stdout.WriteString("true\n")
	// } else {
	// 	os.Stdout.WriteString("false\n")
	// }
}

func atoio(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

// checkpoint
