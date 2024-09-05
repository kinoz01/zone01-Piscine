package main

import (
	"fmt"
)

func main() {
	fmt.Println(WeAreUnique("foo", "boo"))
	fmt.Println(WeAreUnique("", ""))
	fmt.Println(WeAreUnique("abc", "def"))
	fmt.Println(WeAreUnique("26235e50", "4478q92"))
}

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 ==""{
		return -1
	}

	res := []rune{}
	for _, char := range str1 {
		if StrContainsChar(char, str2) {
			continue
		} else if !StrContainsChar(char, string(res)) {
			res = append(res, char)
		}
	}
	for _, char := range str2 {
		if StrContainsChar(char, str1) {
			continue
		} else if !StrContainsChar(char, string(res)) {
			res = append(res, char)
		}
	}
	return len(res)
}

func StrContainsChar(r rune, s string) bool {
	for _, char := range s {
		if char == r {
			return true
		}
	}
	return false
}
