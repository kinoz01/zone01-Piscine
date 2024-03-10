package main

import "fmt"

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Ŝalut!", "alu"))
	fmt.Println(Index(" ", ""))
	fmt.Println(Index("ahgdfalut!alu", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
	fmt.Println(Index("Ollla!", "l"))
	fmt.Println(Index("Nnnxj\\g90X;.u", "nxj\\g90X;"))
}

func Index(s string, toFind string) int {
	runeS := []rune(s)
	runeF := []rune(toFind)
	if len(runeF) > len(runeS) || len(runeS) == 0 || len(runeF) == 0 {
		return -1 // Substring cannot be found if it's longer than the main string
	}
	for i := 0; i < len(runeS)-len(runeF); i++ {
		if string(runeS[i:i+len(runeF)]) == toFind {
			return i // Substring found at index i
		}
	}
	return -1 // Substring not found
}

// quest 5
/***************/
// you can put instead of i<=len(s)-len(toFind)
// a condition
// if i+len(toFind)== (or you can put >=)len(s)-1{
// 	break
// }
// // *********/
