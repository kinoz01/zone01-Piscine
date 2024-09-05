package main

import "os"

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]
	if len(s1) > len(s2) || s1 == "" || s2 == "" {
		return
	}
	j := 0
	for i := 0; i < len(s1); i++ {
		found := false
		for j < len(s2) {
			if s1[i] == s2[j] {
				j++
				found = true
				break
			}
			j++
		}
		if !found {
			return
		}
	}
	os.Stdout.WriteString(s1 + "\n")
}
