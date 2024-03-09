package main

import "fmt"

func Compare(a, b string) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func main() {
	fmt.Println(Compare("Hello!", "Hello!"))
	fmt.Println(Compare("Salut!", "lut!"))
	fmt.Println(Compare("Ola!", "Ol"))
}

// quest 5
// uppercase letters come first
// dictionary comparison so 10 comes before 2
// "10"<"2" in strings logic
