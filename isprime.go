package main

import "fmt"

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(17))
}

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// quest 4
