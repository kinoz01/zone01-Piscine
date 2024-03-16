package main

import "fmt"

func power(n int, p int) int {
	np := 1
	for i := 0; i < p; i++ {
		np = np * n
	}
	return np
}

func convertBaseto10(nbr, base string) int {
	sliceB := []rune(base)
	sliceNbr := []rune(nbr)
	x := len(sliceB)
	index := []int{}
	var n int
	for i := 0; i < len(sliceNbr); i++ {
		for j := 0; j < len(sliceB); j++ {
			if sliceNbr[i] == sliceB[j] {
				index = append(index, j)
			}
		}
	}
	for i := 0; i < len(index); i++ {
		n = n + index[len(index)-1-i]*power(x, i)
	}
	return n
}

func convertBase10toBase(n int, toBase string) string {
	x := len(toBase)
	index := []int{}
	slicetoB := []rune(toBase)
	result := []rune{}
	for n > 0 {
		index = append(index, n%x)
		n /= x
	}
	for i := len(index) - 1; i >= 0; i-- {
		result = append(result, slicetoB[index[i]])
	}
	return string(result)
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return convertBase10toBase(convertBaseto10(nbr, baseFrom), baseTo)
}

func main() {
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
	fmt.Println(ConvertBase("uoi", "choumi", "01"))
	fmt.Println(ConvertBase("1111101", "01", "123"))
	// fmt.Println(power(10,3))
	// fmt.Println(convertBaseto10("uoi", "choumi"))
	// fmt.Println(convertBaseto10("7D", "0123456789ABCDEF"))
	// fmt.Println(convertBase10toBase(125, "0123456789ABCDEF"))
}

// quest 7
