package main

import "os"

func main() {
	if len(os.Args[1:]) != 1 {
		os.Stdout.WriteString("\n")
		return
	}
	s := []rune(os.Args[1])
	r := []rune{}
	t := true
	for i, char := range s {
		if char != ' ' {
			r = append(r, char)
			s[i] = ' '
			t = false
		} else if !t {
			break
		}
	}
	r2 := []rune{}
	f := true
	for _, char := range s {
		if char != ' ' {
			r2 = append(r2, char)
			f = false
		} else if !f {
			r2 = append(r2, char)
			f = true
		}
	}
	if len(r2) == 0 {
		os.Stdout.WriteString(string(r) + "\n")
	} else {
		os.Stdout.WriteString(string(r2) + " " + string(r) + "\n")
	}
}

// checkpoint
