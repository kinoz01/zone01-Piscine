package main

import "os"

func at(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

func it(n int) string {
	var r string
	for n > 0 {
		r = string(n%10+'0') + r
		n /= 10
	}
	return r
}

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
	primes := []int{}
	n := at(s)
	if n < 2 {
		return
	}
	for i := 2; i <= n; i++ {
		if n%i == 0 && isp(i) {
			primes = append(primes, i)
			n /= i
			i = 1
		}
	}
	for i, p := range primes {
		if i != len(primes)-1 {
			os.Stdout.WriteString(it(p) + "*")
		} else {
			os.Stdout.WriteString(it(p) + "\n")
		}

	}
}

func isp(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// checkpoint
