package main

import "fmt"

func main() {
	fmt.Println(Index2("Hello!", "l"))
	fmt.Println(Index2("Ŝalut!", "alu"))
	fmt.Println(Index2(" ", ""))
	fmt.Println(Index2("ahgdfalut!alu", "alu"))
	fmt.Println(Index2("Ola!", "hOl"))
	fmt.Println(Index2("Ollla!", "l"))
	fmt.Println(Index2("Nnnxj\\g90X;.u", "nxj\\g90X;"))
}

func Index2(s string, toFind string) int {
	sliceS := []rune(s)
	sliceTofind := []rune(toFind)
	count := -1
	found := false 
	if len(s) == 0 || len(toFind) == 0 || len(s) < len(toFind) {
		return -1
	}
	for i := 0; i < len(sliceS); i++ {
		if found == true {
			break
		}
		if sliceTofind[0] == sliceS[i] {
			count = i // check for the first character when its found
		}
		for j := 1; j < len(sliceTofind); j++ {
			if sliceTofind[j] == sliceS[count+j] {
				continue
			} else {
				count = -1
				break
			}
		}
		if count >= 0 {
			found = true // adding the found mechansim to stop searching if the substring is found
		}

	}
	return count
}
