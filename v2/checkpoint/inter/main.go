package main

import "os"

func main() {
	if len(os.Args) != 3 {
		return
	}
	result := []rune{}

	for _, c1 :=range os.Args[1] {
		for _, c2 :=range os.Args[2] {
			if c1 == c2 && !Contains(string(result), c1) {
				result = append(result, c1)
			}
		}
	}
	os.Stdout.WriteString(string(result))
}

func Contains(s string, r rune) bool{
	for _, char := range s {
		if r == char {
			return true
		}
	}
	return false
}
