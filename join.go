package main

import "fmt"

func Join(strs []string, sep string) string {
	sliceString := []string(strs)
	result := strs[0]
	for i := 1; i < len(sliceString); i++ {
		result = result + sep + sliceString[i]
	}
	return result
}

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(Join(toConcat, ":"))
}

// quest 5
