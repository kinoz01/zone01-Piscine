package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}

func CountAlpha(s string) int {
	c := 0
	for _, char := range s {
		if (char >= 'a' && char<= 'z') || (char >= 'A' && char<= 'Z') {
			c++
		}
	}
	return c
}
