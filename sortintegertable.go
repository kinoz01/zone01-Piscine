package main

import "fmt"

func SortIntegerTable(table []int) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func main() {
	s := []int{4, 4, 3, 5, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

// quest 3 (using bubble sorting algo)
