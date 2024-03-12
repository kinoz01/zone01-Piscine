package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	union := args[0] + args[1]
	result := []rune{}
	if len(args) != 2 || len(args[0]) == 0 || len(args[1]) == 0 {
		fmt.Println()
		return
	}
	for _, char := range union { // range over union
		found := false // use found so we don't add similar characters
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

// checkpoint
