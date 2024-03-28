package main

import "fmt"

func AddFront(s string, slice []string) []string {
	r := []string{s}
	for _, str := range slice {
		r = append(r, str)
	}
	return r
}

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}

// checkpoint
