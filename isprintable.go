package main

import "fmt"

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if char < ' ' || char > '~' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrintable("Hello"))
	fmt.Println(IsPrintable("Hello\n"))
}

// quest 5
// https://theasciicode.com.ar/
