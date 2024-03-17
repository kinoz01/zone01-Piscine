package main

import "fmt"

func Unmatch(a []int) int {
	for i := 0; i < len(a); i++ {
		count := 1
		for j := 0; j < len(a); j++ {
			if i != j && a[i] == a[j] {
				count++
			}
		}
		if count%2 != 0 {
			return a[i]
		}
	}
	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4, 4, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch)
}

// hackathon
