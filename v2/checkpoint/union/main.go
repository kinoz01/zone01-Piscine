package main

import "os"

func main() {
	if len(os.Args) != 3 {
		os.Stdout.WriteString("\n")
		return
	}

	s := os.Args[1] + os.Args[2]

	result := ""
	for _, char := range s {
		if !Contains(result, char) {
			result += string(char)
		}
	}
	os.Stdout.WriteString(result+ "\n")
}

func Contains(s string, r rune) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}