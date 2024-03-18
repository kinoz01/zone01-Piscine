package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Stdout.WriteString("0\n")
	} else {
		os.Stdout.WriteString(printn(len(args)) + "\n")
	}
}

func printn(n int) string {
	var r string
	for n > 0 {
		r = string(n%10+'0') + r
		n /= 10
	}
	return r
}

// checkpoint
