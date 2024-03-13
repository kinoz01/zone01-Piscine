package main

import "fmt"

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return "\n"
	}
	result := args[0]
	for i := 1; i < len(args); i++ {
		result = result + "\n" + args[i]
	}
	return result
}

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(ConcatParams(test))
}

// quest 7
