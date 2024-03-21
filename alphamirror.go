package main

import "os"

func main() {
	if len(os.Args) != 2 {
		return
	}
	lo := "abcdefghijklmnopqrstuvwxyz"
	up := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	s := []rune(os.Args[1])
	l := []rune(lo)
	u := []rune(up)
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(l); j++ {
			if s[i] == l[j] {
				s[i] = l[len(l)-1-j]
				break
			}
		}
		for j := 0; j < len(u); j++ {
			if s[i] == u[j] {
				s[i] = u[len(u)-1-j]
				break
			}
		}
	}
	os.Stdout.WriteString(string(s) + "\n")
}


/**************** Another Way*****************/
// func main() {
// 	if len(os.Args) == 2 {
// 		args := []rune(os.Args[1])
// 		for i, ch := range args {
// 			if ch >= 'a' && ch <= 'z' {
// 				args[i] = 'z' - ch + 'a'
// 			}
// 			if ch >= 'A' && ch <= 'Z' {
// 				args[i] = 'Z' - ch + 'A'
// 			}
// 		}
// 		os.Stdout.WriteString(string(args))
// 	}
// 	os.Stdout.WriteString("\n")
// }
/********************************************/

// checkpoint
