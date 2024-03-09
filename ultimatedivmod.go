package main

import "fmt"

func UltimateDivMod(a *int, b *int) {
	c := *a // temporair value to save a value (*a)
	*a = *a / *b
	*b = c % *b
}

func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

// Quest 3
