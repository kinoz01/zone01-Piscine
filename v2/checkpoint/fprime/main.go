package main

import "os"

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	n := Atoi(os.Args[1])
	if n <2 {
		return
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if n%i == 0 && IsPrime(i) {
			primes = append(primes, i)
			n /= i
			i = 1
		}
	}
	result := ""
	for _, p := range primes {
		result += Itoa(p) + "*"
	}
	os.Stdout.WriteString(result[:len(result)-1] + "\n")
}

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Atoi(s string) (n int) {
	neg := false
	for i, char := range s {
		if i == 0 && char == '-' {
			neg = true
			continue
		}
		if char > '9' || char < '0' {
			return 0
		}
		n = n*10 + int(char-'0')
	}
	if neg {
		return -n
	} else {
		return n
	}
}

func Itoa(n int) (s string) {
	var sign string
	if n < 0 {
		sign = "-"
		n = -n
	}
	for n > 0 {
		s = string(rune(n%10)+'0') + s
		n /= 10
	}
	return sign + s
}
