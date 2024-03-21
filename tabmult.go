package main

import "os"

func ituoi(n int) string {
	var r string
	var negative bool
	if n == 0 {
		return "0"
	}
	if n < 0 {
		n = -n
		negative = true
	}
	for n > 0 {
		r = string(n%10+'0') + r
		n /= 10
	}
	if negative {
		return "-" + r
	}
	return r
}

func atuoi(str string) int {
	var n int
	s := []rune(str)
	first := true
	for i := 0; i < len(s); i++ {
		if (s[0] == '-' || s[0] == '+') && first {
			first = false
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			os.Stdout.WriteString("\n")
			os.Exit(0)
		}
		n = n*10 + int(s[i]-'0')
	}
	if s[0] == '-' {
		return -n
	}
	return n
}

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}
	num := os.Args[1]
	if len(num) == 0 {
		os.Stdout.WriteString("\n")
		return
	}
	c := 1
	for i := '1'; i <= '9'; i++ {
		os.Stdout.WriteString(string(i) + " x " + num + " = " + ituoi(atuoi(num)*c) + "\n")
		c++
	}
}

// checkpoint / atoi / itoa
