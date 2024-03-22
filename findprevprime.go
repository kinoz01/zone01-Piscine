package main

import "fmt"

func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	for nb >= 2 {
		if isprime(nb) {
			return nb
		}
		nb -= 1
	}
	return 0
}

func isprime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(FindPrevPrime(10))
	fmt.Println(FindPrevPrime(42))
}

// checkpoint
