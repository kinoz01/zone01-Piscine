package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(WordFlip("First second last"))
	fmt.Print(WordFlip(""))
	fmt.Print(WordFlip("     "))
	fmt.Print(WordFlip(" hello  all  of  you! "))
}

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output"
	}
	a := strings.Fields(str)
	result := ""
	if len(a) < 1 {
		return "\n"
	}
	for i := len(a) - 1; i >= 0; i-- {
		result += a[i] + " "
	}
	return result[:len(result)-1] + "\n"
}
