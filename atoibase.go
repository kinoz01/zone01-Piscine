package main

import "fmt"

func AtoiBase(s string, base string) int {
	sliceBase := []rune(base)
	if len(sliceBase) < 2 {
		return 0
	}
	for i := 0; i < len(sliceBase); i++ {
		for j := 0; j < len(sliceBase); j++ {
			if i != j && sliceBase[i] == sliceBase[j] {
				return 0
			}
		}
		if sliceBase[i] == '+' || sliceBase[i] == '-' {
			return 0
		}
	}
	x := len(base)
	sliceS := []rune(s)
	var n int
	index := []int{}
	for i := len(sliceS) - 1; i >= 0; i-- {
		for j := 0; j < len(sliceBase); j++ {
			if sliceS[i] == sliceBase[j] {
				index = append(index, j)
			}
		}
	}
	//fmt.Println(index)
	for i := 0; i < len(index); i++ {
		n = n + index[i]*pp(x, i)
	}
	return n
}

func pp(n, p int) int {
	m := 1
	for i := 0; i < p; i++ {
		m = m * n
	}
	return m
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

// checkpoint
