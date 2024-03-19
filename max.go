package main

import "fmt"

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	max := a[0]
	for _, n := range a {
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	a := []int{223, 123, 1, 11, 55, 910}
	max := Max(a)
	fmt.Println(max)
}

// checkpoint
