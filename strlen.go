package main

import "fmt"

func StrLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}

func main() {
	l := StrLen("Hello World!")
	n := StrLen("He!îo")
	fmt.Println(l)
	fmt.Println(n)
}

// Quest 3 / checkpoint
/***************** another way will be *************/
// func StrLen(s string) int {
// 	runeS:=[]rune(s)
// 	return len(runeS)
// }

/***************** why return len(s) won't work for "He!îo" **************************/
/*The len(s) function in Go returns the number of bytes,
not the number of characters, in a string. Go strings are
UTF-8 encoded, and UTF-8 is a variable-width encoding,
 meaning different characters can use a different number
 of bytes.

In the string "He!îo", most of the characters are
ASCII and occupy 1 byte each. However, the character î
is not an ASCII character and is represented in UTF-8
using two bytes (0xC3 0xAE) */

