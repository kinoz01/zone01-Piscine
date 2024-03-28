package main

import "os"

func main() {
	if len(os.Args[1:]) != 1 || os.Args[1] == "" {
		return
	}
	s := os.Args[1]
	r := []rune{}
	for _, char := range s {
		if char != ' ' {
			r = append(r, char)
		} else if len(r) != 0 {
			break
		}
	}
	os.Stdout.WriteString(string(r) + "\n")
}

// checkpoint
