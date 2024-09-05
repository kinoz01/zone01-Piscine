package main

import (
	"fmt"
	//"checkpoint/slice/solution"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", solutions.Slice(a, 1))
	fmt.Printf("%#v\n", solutions.Slice(a, 1, 10))
	fmt.Printf("%#v\n", solutions.Slice(a, -3))
	fmt.Printf("%#v\n", solutions.Slice(a, -2, 1))
	fmt.Printf("%#v\n", solutions.Slice(a, 2, 0))
}

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return a
	}

	first := nbrs[0]
	if len(nbrs) == 1 {
		if first < 0 {
			first = len(a) + first
			if first < 0 {
				return a
			}
		}
		return a[first:]
	}
	second := nbrs[1]

	first = isNegative(a, first)
	second = isNegative(a, second)

	if first > second {
		return nil
	}
	return a[first:second]
}

func isNegative(a []string, n int) int {
	if n < 0 {
		n = len(a) + n
	}
	if n < 0 {
		n = 0
	} else if n > len(a) {
		n = len(a)
	}
	return n

}
