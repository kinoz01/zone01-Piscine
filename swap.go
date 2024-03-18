package main

import "fmt"

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

// quest 3 / chekpoint
