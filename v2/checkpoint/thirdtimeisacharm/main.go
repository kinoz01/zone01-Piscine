package main

import "fmt"

func main() {
	fmt.Print(ThirdTimeIsACharm("123456789"))
	fmt.Print(ThirdTimeIsACharm(""))
	fmt.Print(ThirdTimeIsACharm("a b c d e f"))
	fmt.Print(ThirdTimeIsACharm("12"))
}

func ThirdTimeIsACharm(str string) string {
	var r []rune
	for i, char := range str {
		if (i-2)%3 == 0 {
			r = append(r, char)
		}
	}
	return string(r) + "\n"
}
