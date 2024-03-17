package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}

// quest 2

// go mod init go.mod
// go mod tidy

// git pull origin master

/* A rune represents a Unicode code point,
which means it can encode any Unicode character,
not just ASCII characters. Since Unicode aims to
include all character sets, this makes runes capable
of representing virtually any character from any
language.*/
//
// https://theasciicode.com.ar/
// 

