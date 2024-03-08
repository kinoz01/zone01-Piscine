package main

import (
	"fmt"
	"strconv"
)

func fromto(from int, to int) string {

	var result string

	if from <= 0 || to <= 0 || from > 99 || to > 99 {
		return "invalid\n"
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
	fmt.Print(fromto(1, 10))
	fmt.Print(fromto(10, 1))
	fmt.Print(fromto(1, 20))
	fmt.Print(fromto(10, 10))
	fmt.Print(fromto(11, 13))
	fmt.Print(fromto(9, 13))
	fmt.Print(fromto(15, 5))
	fmt.Print(fromto(-1, 10))
}

// checkpoint
