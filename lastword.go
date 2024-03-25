package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	sliceA := []rune(args[0])
	r := []rune{}
	for i := len(sliceA) - 1; i >= 0; i-- {
		if sliceA[i] != ' ' {
			r = append(r, sliceA[i])
		} else if len(r) != 0 { // if sliceA[i] == ' ' and r is not empty
			break
		}
	}
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	if len(r) != 0 {
		os.Stdout.WriteString(string(r) + "\n")
	} else {
		return
	}
}

// checkpoint
