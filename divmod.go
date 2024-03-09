package main

import "fmt"

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

func main() {
	a := 13
	b := 2
	var div int
	var mod int
	DivMod(a, b, &div, &mod)
	fmt.Println(div)
	fmt.Println(mod)
}

// Quest 3

/* an important this about pointers is changing their
values in a function a will change their values in all
other functions when they are being used*/
