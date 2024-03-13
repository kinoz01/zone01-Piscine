package main

import "fmt"

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	intSlice := make([]int, max-min)
	for i := min; i < max; i++ {
		intSlice[i-min] = i
	}
	return intSlice
}

func main() {
	fmt.Println(MakeRange(5, 10))
	fmt.Println(MakeRange(10, 5))
}

// quest 7
