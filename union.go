package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	union := args[0] + args[1]
	var result []rune
	found := false
	if len(args) != 2 {
		fmt.Println()
		return
	}
	for _, char := range union {
		found = false
		for _, n := range result {
			if char == n {
				found = true
				break
			}
		}
		if found == false {
			result = append(result, char)
		}
	}
	fmt.Println(string(result))
}
