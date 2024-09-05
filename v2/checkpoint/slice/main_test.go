package main

import (
	"checkpoint/slice/solution" // Replace with the correct package path
	"fmt"
	"strconv"
	"testing"

	"github.com/01-edu/go-tests/lib/random" // Assuming the random package is in the same directory or accessible in your GOPATH
	"github.com/01-edu/go-tests/lib/chars"
)

// ANSI escape codes for coloring and formatting
const (
	ColorReset     = "\033[0m"
	ColorRed       = "\033[31m"
	ColorGreen     = "\033[32m"
	ColorYellow    = "\033[33m"
	ColorCyan      = "\033[36m"
	ColorBlue      = "\033[34m"
	ColorBold      = "\033[1m"
	ColorUnderline = "\033[4m"
)

func TestSlice(t *testing.T) {
	// Predefined test cases
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	// Random test cases
	s := random.StrSlice(chars.Words)
	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	fmt.Printf("%sRunning Slice tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, elem := range elems {
		a := elem[0].([]string)
		nbr := []int{}
		for _, v := range elem[1:] {
			nbr = append(nbr, v.(int))
		}

		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the Slice function in solution/main.go to get the expected result
			expectedResult := solution.Slice(a, nbr...)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {a: %v, nbr: %v}%s\n", ColorBlue, i+1, a, nbr, ColorReset)
			fmt.Printf("%sExpected Output: %v%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the Slice function in main.go
			mainResult := Slice(a, nbr...)
			fmt.Printf("%sYour Output: %v%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if fmt.Sprintf("%v", mainResult) != fmt.Sprintf("%v", expectedResult) {
				fmt.Printf("%s%sMismatch for input {a: %v, nbr: %v}%s\n", ColorRed, ColorBold, a, nbr, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour Slice() = %v\nExpected = %v", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
