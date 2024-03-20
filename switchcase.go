package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	runeA := []rune(args[0])
	for i := 0; i < len(runeA); i++ {
		if runeA[i] >= 'a' && runeA[i] <= 'z' {
			runeA[i] -= 32
		} else if runeA[i] >= 'A' && runeA[i] <= 'Z' {
			runeA[i] += 32
		}
	}
	os.Stdout.WriteString(string(runeA) + "\n")
}

// checkpoint
