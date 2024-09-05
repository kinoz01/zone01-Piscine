package main

import (
	"fmt"
)

func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Print(FifthAndSkip("This is a short sentence"))
	fmt.Print(FifthAndSkip("1234"))
}

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Input\n"
	}
	str = RemoveSpace(str)
	var result string
	
	for i, char := range []rune(str) {
		if (i-5)%6 == 0 && i!=0{
			result += " "
			continue
		}
		result += string(char)
		
	}
	return result+"\n"
}

func RemoveSpace(s string) string {
	r := ""
	for _, char := range s {
		if char == ' ' {
			continue
		}
		r += string(char)
	}
	return r
}
