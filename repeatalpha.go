package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	s := args[0]
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			for i := 'a'; i <= char; i++ {
				os.Stdout.WriteString(string(char))
			}
		} else if char >= 'A' && char <= 'Z' {
			for i := 'A'; i <= char; i++ {
				os.Stdout.WriteString(string(char))
			}
		} else {
			os.Stdout.WriteString(string(char))
		}
	}
	os.Stdout.WriteString("\n")
}

// checkpoint
