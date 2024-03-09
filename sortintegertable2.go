package main

import "fmt"

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				table[i], table[j] = table[j], table[i]
			}
		}
	}
}

func main() {
	s := []int{4, 4, 3, 5, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

// quest 3 (selection sort algorithm)
