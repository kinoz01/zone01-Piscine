package main

import (
	"fmt"
	"strconv"
)

func FromTo(from int, to int) string {

	var result string

	if from < 0 || to < 0 || from > 99 || to > 99 {
		return "Invalid\n"
	}
	if from > to {
		for i := from; i >= to; i-- {
			if i < 10 {
				result += "0" + strconv.Itoa(i)
			} else {
				result += strconv.Itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	}
	if from < to {
		for i := from; i <= to; i++ {
			if i < 10 {
				result = result + "0" + strconv.Itoa(i)
			} else {
				result = result + strconv.Itoa(i)
			}
			if i != to {
				result += ", "
			}
		}
	}
	if from == to {
		result = strconv.Itoa(to)
	}
	return result + "\n"
}

func main() {
	fmt.Print(FromTo(1, 10))
	fmt.Print(FromTo(10, 1))
	fmt.Print(FromTo(1, 20))
	fmt.Print(FromTo(10, 10))
	fmt.Print(FromTo(11, 13))
	fmt.Print(FromTo(9, 13))
	fmt.Print(FromTo(15, 5))
	fmt.Print(FromTo(-1, 10))
}
