package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	if !CheckIfCamel(s) {
		return s
	}

	return CamelToSnake(s)
}

func CheckIfCamel(s string) bool {
	r := []rune(s)
	if r[len(r)-1] >= 'A' && r[len(r)-1] <= 'Z' {
		return false
	}
	for i, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
		if i+1 < len(s) && (char >= 'A' && char <= 'Z') && (s[i+1] >= 'A' && s[i+1] <= 'Z') {
			return false
		}
	}
	return true
}

func CamelToSnake(s string) string {
	rSlice := []rune{}
	r := []rune(s)
	for i, char := range r {
		rSlice = append(rSlice, char)
		if i+1 < len(r) && r[i+1] >= 'A' && r[i+1] <= 'Z' {
			rSlice = append(rSlice, '_')
		}
	}
	return string(rSlice)
}
