package main

import "fmt"

func AppendRange(min, max int) []int {
	intSlice := []int{}
	if min >= max {
		return intSlice
	}
	for i := min; i < max; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}

func main() {
	fmt.Println(AppendRange(5, 10))
	fmt.Println(AppendRange(10, 5))
}

// quest 7
