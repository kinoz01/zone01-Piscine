package main

import "fmt"

func TrimAtoi(s string) int {
	sliceS := []rune(s)
	index := []int{}
	t := 1
	var n int
	for i, char := range s {
		if char >= '0' && char <= '9' {
			index = append(index, i)
		}
		if len(index) == 0 && char == '-' {
			t = -1
		}
	}
	if len(index) == 0 {
		return 0
	}
	for i := 0; i < len(index); i++ {
		n = n*10 + int(sliceS[index[i]]-'0')
	}
	return n * t
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

// quest 5 / checkpoint
