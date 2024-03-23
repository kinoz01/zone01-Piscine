package main

import "fmt"

func LoafOfBread(str string) string {
	s := []rune(str)
	r := []rune{}
	if len(s) < 5 {
		return "Invalid Output\n"
	}
	var stp bool // if we don't use stp we will enter the second if which something we don't want untill we append a new "alpha"
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			r = append(r, s[i])
			stp = true
			count++ // we count each "alpha" appended
		}
		if count%5 == 0 && stp && i+1 < len(s) && len(r) != 0 {
			if s[i+1] == ' ' {
				stp = false
			} else {
				s[i+1] = ' ' // replace the character with space
				stp = false
			}
			r = append(r, ' ')
		}
	}
	return string(r) + "\n"
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}

// hackathon
