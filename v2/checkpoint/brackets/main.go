package main

import "os"

func main() {
	if len(os.Args) <= 1 {
		return
	}

	for _, str := range os.Args[1:] {
		if CheckBrackets(str) {
			os.Stdout.WriteString("OK\n")
		} else {
			os.Stdout.WriteString("Error\n")
		}
	}

}

func CheckBrackets(s string) bool {
	bra := map[rune]rune {
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []rune
	for _, char := range s {
		switch char{
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 || stack[len(stack)-1] != bra[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
