package main

import "fmt"

func Chunk(slice []int, size int) {
	if size <= 0 {
		fmt.Println()
		return
	}
	start := 0
	result := [][]int{}
	for i := 1; i < len(slice); i++ {
		if i%size == 0 {
			result = append(result, slice[start:i])
			start = i
		}
	}
	result = append(result, slice[start:])
	fmt.Println(result)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

// checkpoint // strchunks revisited
