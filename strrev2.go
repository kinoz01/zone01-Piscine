package main

import "fmt"

func StrRev2(s string) string {
	runeS := []rune(s)
	for i, j := 0, len(runeS)-1; i < j; i, j = i+1, j-1 {
		runeS[i], runeS[j] = runeS[j], runeS[i]
	}
	return string(runeS)
}

func main() {
	s := "Helloô World!"  // using this make obvious that we need to work with len(runeS) instead of len(s)
	s = StrRev2(s)
	fmt.Println(s)
}

// quest 3 (using swap and loop)

// You can also use:
/*for i := 0; i < len(s)/2; i++ {
		j := len(runeS) - i - 1
		runeS[i], runeS[j] = runeS[j], runeS[i]
	}*/
