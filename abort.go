package main

import "fmt"

func Abort(a, b, c, d, e int) int {
	return (a+b+c+d+e)/5
}

func main() {
	middle := Abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}

// hackathon
