package main

import "fmt"

func StrRev(s string) string {
	runeS := []rune(s)
	revers := make([]rune, len(runeS)) // if you don't use make here to define len you will get out of range error
	for i := 0; i < len(runeS); i++ {
		revers[i] = runeS[len(runeS)-i-1]
	}
	return string(revers) // convert slice of runes to a string
}

func main() {
	s := "Helloô World!"
	s = StrRev(s)
	fmt.Println(s)
}

// quest 3 (using reverse array algorithm)
// IMPORTANT: len(s) != len(runeS) you have to use the last to count correctly
