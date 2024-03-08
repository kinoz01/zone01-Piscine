package main

import "fmt"

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	a := 0
	b := &a
	n := &b
	UltimatePointOne(&n)
	fmt.Println(a)
}

// Quest 3
