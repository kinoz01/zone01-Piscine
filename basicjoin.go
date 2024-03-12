package main

import "fmt"

func BasicJoin(elems []string) string {
	sliceString := []string(elems)
	var result string
	for i := 0; i < len(sliceString); i++ {
		result = result + sliceString[i]
	}
	return result
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}

// quest 5
