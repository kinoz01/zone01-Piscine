package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	s1 := []rune(rmvRepStr(args[0]))
	s2 := []rune(args[1])
	r := []rune{}
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s2[j] == s1[i] {
				r = append(r, s1[i])
				break
			}
		}
	}
	os.Stdout.WriteString(string(r) + "\n")
}

func rmvRepStr(s string) string {
	r := []rune{}
	sliceS := []rune(s)
	for i := 0; i < len(sliceS); i++ {
		rep := false
		for j := 0; j < len(r); j++ {
			if sliceS[i] == r[j] {
				rep = true
				continue
			}
		}
		if !rep {
			r = append(r, sliceS[i])
		}
	}
	return string(r)
}

// checkpoint

/********* in case where chaaracters have to match the
exact dispaly order then u have to use: */
// 	j := c // we gonna start our loop (searching) from where we found the last character (c=j+1)
// 	for j < len(s2) {
// 		if s1[i] == s2[j] {
// 			r = append (r, s1[i])
// 			c = j + 1
// 			break
// 		} else {
// 			j++
// 		}
// 	}
// }
// (like we did in wdmatch)
