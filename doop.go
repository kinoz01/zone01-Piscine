package main

import "os"

// import "fmt"

const maxInt = int(^uint(0) >> 1)
const minInt = -maxInt - 1

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	result := doop(arg[0], arg[1], arg[2])
	os.Stdout.WriteString(result)
}

func doop(s0, o, s2 string) string {
	var modulo int
	var div int
	var plus int
	var moin int
	var product int
	first := true
	first2 := true

	if o != "+" && o != "-" && o != "/" && o != "*" && o != "%" {
		return ""
	}
	for i := 0; i < len(s0); i++ {
		if (s0[0] == '+' || s0[0] == '-') && first {
			first = false
			continue
		}
		if s0[i] > '9' || s0[i] < '0' {

			return ""
		}
	}
	for i := 0; i < len(s2); i++ {
		if (s2[0] == '+' || s2[0] == '-') && first2 {
			first2 = false
			continue
		}
		if s2[i] > '9' || s2[i] < '0' {

			return ""
		}
	}
	n1 := atoii(s0)
	n2 := atoii(s2)
	if o == "%" {
		if n2 == 0 {
			return "No modulo by 0"
		}
		modulo = n1 % n2
		return itoaa(modulo)
	}
	if o == "/" {
		if n2 == 0 {
			return "No division by 0"
		}
		div = n1 / n2
		return itoaa(div)+"\n"
	}
	if o == "+" {
		if addCheck(n1, n2) {
			return ""
		}
		plus = n1 + n2
		return itoaa(plus)+"\n"
	}
	if o == "*" {
		if multipCheck(n1, n2) {
			return ""
		}
		if n2 == 0 || n1 == 0 {
			return "0"
		}
		product = n1 * n2
		return itoaa(product)+"\n"
	}
	if o == "-" {
		if subCheck(n1, n2) {
			return ""
		}
		moin = n1 - n2
		return itoaa(moin)+"\n"
	}
	return ""
}

func itoaa(nbr int) string {
	var r string
	var n int
	if nbr == 0 {
		return "0"
	}
	if nbr < 0 {
		n = -nbr
	} else {
		n = nbr
	}
	for n > 0 {
		r = string(n%10+'0') + r
		n /= 10
	}
	if nbr < 0 {
		return "-" + r
	}
	return r
}

func atoii(s string) int {
	var n int
	sign := 1
	for _, digit := range s {
		if digit == '+' {
			continue
		}
		if digit == '-' {
			sign = -1
			continue
		}
		n = n*10 + int(digit-'0')
	}
	return n * sign
}

func addCheck(a, b int) bool {
	if a > 0 && b > 0 && a + b > maxInt{
		return true // overflow
	}
	if a < 0 && b < 0 && a - b < minInt{
		return true // underflow
	}
	return false // no overflow
}

func subCheck(a, b int) bool {
	if a < 0 && b > 0 && minInt+a < -b {
		return true // Underflow
	}
	if a > 0 && b < 0 && maxInt-a < -b {
		return true // Overflow
	}
	return false
}

func multipCheck(a, b int) bool {
	if a > 0 && b > 0 && a * b > maxInt {
		return true // Overflow
	}
	if a < 0 && b < 0 && a * b < maxInt {
		return true // Overflow
	}
	if a < 0 && b > 0 && a * b < minInt {
		return true // Underflow
	}
	if a > 0 && b < 0 && b * a < minInt {
		return true // Underflow
	}
	return false
}
