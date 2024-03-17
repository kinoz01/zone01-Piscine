package main

import "fmt"

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i+1 < len(a); i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}

func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	} else {
		return 0
	}
}

func main() {
	a1 := []int{0, 1, 2, 3, 5, 7}
	a2 := []int{0, 2, 1, 3}

	result1 := IsSorted(f, a1)
	result2 := IsSorted(f, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}

// quest 9
