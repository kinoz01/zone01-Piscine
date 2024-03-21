package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}
	s := []rune(args[0])
	for i := 0; i < len(s); i++ {
		if s[i] == rune(args[1][0]) {
			s[i] = rune(args[2][0])
		}
	}
	os.Stdout.WriteString(string(s) + "\n")
}

// checkpoint
