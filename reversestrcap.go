package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	for _, str := range args {
		for i := 0; i < len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' && ((i+1 < len(str) && str[i+1] == ' ') || (i == len(str)-1)) { // the condition str[i] >= 'a' && str[i] <= 'z' is factorized
				os.Stdout.WriteString(string(str[i] - 32))
			} else if str[i] >= 'A' && str[i] <= 'Z' && i+1 < len(str) && str[i+1] != ' ' && i != len(str)-1 {
				os.Stdout.WriteString(string(str[i] + 32))
			} else {
				os.Stdout.WriteString(string(str[i]))
			}
		}
		os.Stdout.WriteString("\n")
	}
}

// checkpoint
