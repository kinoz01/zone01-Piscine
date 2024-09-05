package main

import (
	"fmt"
)

func main() {
	fmt.Println(SaveAndMiss("123456789", 3))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))
	fmt.Println(SaveAndMiss("e 5£ @ 8* 7 =56 ;", 2))
}

func SaveAndMiss(arg string, num int) string {
	if num <= 0 || num > len(arg) {
		return arg
	}

	str := ""
	for i :=0; i<len(arg); i++ {
		if i!=0 && i%num == 0 {
			i += num
			if i>len(arg){
				break
			}
		}
		str += string(rune(arg[i]))
	}
	return str
}
