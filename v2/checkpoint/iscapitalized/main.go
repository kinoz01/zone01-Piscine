package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func IsCapitalized(s string) bool {
	if s == "" {
		return false
	}
	words := Fields(s, ' ')
	for _, str := range words {
		if str[0] >= 'a' && str[0] <= 'z' {
			return false
		}
	}
	return true
}

func Fields(s string, sep rune) (res []string) {

	if s == "" {
		return nil
	}
	var stack []rune
	for _, r := range s {
		if r == sep && len(stack) > 0 {
			res = append(res, string(stack))
			stack = nil
			continue
		}
		stack = append(stack, r)
	}
	if len(stack)>0 {
		res = append(res, string(stack))
	}
	
	return res
}
