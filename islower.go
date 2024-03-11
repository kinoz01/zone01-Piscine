package main

import "fmt"

func IsLower(s string) bool {
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] > 'z' || sliceS[i] < 'a' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}

// quest 5
