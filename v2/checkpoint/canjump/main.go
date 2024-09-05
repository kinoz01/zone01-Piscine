package main

import (
	"checkpoint/canjump/solution"
	"fmt"
)

func main() {
	fmt.Println(CanJump([]uint{0}))
	fmt.Println(solution.CanJump([]uint{0}))
}

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}
	MaxReach := 0
	for i, n := range steps {
		if n != 0 {
			MaxReach = i + int(n)
		}
		if MaxReach == len(steps)-1 || len(steps) == 1 {
			return true
		}
	}
	return false
}
