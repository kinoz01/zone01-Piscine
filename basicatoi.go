package main

import "fmt"

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

// quest 3

/*
int('0')=48 -----> int('1'-'0') = int(49-48) = int(1) = 1
In Go, when you convert a character constant like '0'
to an integer using int('0'), you're getting the ASCII
(or Unicode code point) value of that character.
The ASCII value of '0' is 48. So, int('0') will return 48.*/
