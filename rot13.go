package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	sliceS := []rune(args[0])
	for i := 0; i < len(sliceS); i++ {
		if sliceS[i] >= 'a' && sliceS[i] <= 'm' || sliceS[i] >= 'A' && sliceS[i] <= 'M' {
			sliceS[i] += 13
		} else if sliceS[i] > 'm' && sliceS[i] <= 'z' || sliceS[i] > 'M' && sliceS[i] <= 'Z' {
			sliceS[i] -= 13
		}
	}
	os.Stdout.WriteString(string(sliceS) + "\n")
}

// checkpoint
