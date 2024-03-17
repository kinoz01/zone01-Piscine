package main

import "fmt"

func Map(f func(int) bool, a []int) []bool {
	boolslice := make([]bool, len(a))
	for i := 0; i < len(a); i++ {
		if f(a[i]) {
			boolslice[i] = true
		} else {
			boolslice[i] = false
		}
	}
	return boolslice
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	result := Map(IsPrime, a)
	fmt.Println(result)
}

// quest 9
