package main

import "os"

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
	if len(args) != 1 {
		return
	}
	n := ato(args[0])
	var r string
	if exit || n == 0 {
		os.Stdout.WriteString("00000000\n")
		return
	}
	if n > 0 {
		for n > 0 {
			r = string(n%2+'0') + r
			n /= 2
		}
		for i := 0; i < 8-len(r); i++ {
			os.Stdout.WriteString("0")
		}
		os.Stdout.WriteString(r + "\n")
	}
	if n < 0 {
		return // I don't know how we should handle negatives
	}
}

// checkpoint
