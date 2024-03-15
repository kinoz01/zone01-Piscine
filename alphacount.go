package main

import "fmt"

func main() {
	s := "Hello 78 World!    4455 /"
	nb := AlphaCount(s)
	fmt.Println(nb)
}

func AlphaCount(s string) int {
	count := 0
	runeS := []rune(s)
	for i := 0; i < len(runeS); i++ {
		if runeS[i] >= 'A' && runeS[i] <= 'Z' {
			count++
		}
		if runeS[i] >= 'a' && runeS[i] <= 'z' {
			count++
		}
	}
	return count
}

// quest 5
