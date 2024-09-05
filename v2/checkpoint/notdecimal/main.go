package main

import (
	"fmt"
)

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}
	if quit := CheckNum(dec); quit {
		return dec +"\n"
	}
	num := ""
	for _, char := range dec {
		if char == '.' {
			continue
		}
		num += string(char)
	}
	return Itoa(Atoi(num))+"\n"
}

func Atoi(s string)int {
	var n int
	sliceS := []rune(s)
	first := true
	for i := 0; i < len(sliceS); i++ {
		if (sliceS[0] == '-' || sliceS[0] == '+') && first {
			first = false
			continue
		}
		if sliceS[i] > '9' || sliceS[i] < '0' {
			return 0
		}
		n = n*10 + int(sliceS[i]-'0')
	}
	if sliceS[0] == '-' {
		return -n
	}
	return n
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

func CheckNum(s string) bool {
	if s[0] == '-' {
		s = s[1:]
	}
	if s[0] == '.' {
		return true
	}
	ptFound := false
	for _, char := range s {
		if char == '.' && !ptFound{
			ptFound =true
			continue
		}
		if char < '0' || char > '9' {
			return true
		}
	}
	return false
}
