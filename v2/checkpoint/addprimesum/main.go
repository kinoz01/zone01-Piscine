package main

import "os"

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		os.Stdout.WriteString("0\n")
		return
	}
	n := Atoi(os.Args[1])
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}
	res :=0
	for n>0 {
		if isPrime(n) {
			res += n
		}
		n--
	}
	os.Stdout.WriteString(Itoa(res)+"\n")
}

func Atoi(s string) int {
	var n int
	for _, char := range s {
		if char > '9' || char < '0' {
			return 0
		}
		n = n*10 + int(char-'0')
	}
	return n
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i:=2; i*i<=n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Itoa(n int)string{
	var s string
	for n>0 {
		s = string(rune(n%10)+'0') +s 
		n /= 10
	}
	return s
}
