package main

import "fmt"

func main() {
	fmt.Println(FindNextPrime(9))
	fmt.Println(FindNextPrime(4))
	fmt.Println(FindNextPrime(21))
}

func FindNextPrime(nb int) int {
	prime := nb
	if nb <= 1 {
		return 2
	}
	for i := 0; !isprime(prime); i++ {
		prime = nb + i
	}
	return prime
}

func isprime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// quest 4
