package main

import (
	"checkpoint/findprevprime/solution" // Replace with the correct package path
	"fmt"
	"strconv"
	"testing"

	"github.com/01-edu/go-tests/lib/random" // Assuming the random package is in the same directory or accessible in your GOPATH
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

func TestFindPrevPrime(t *testing.T) {
	// Predefined test cases
	args := []int{
		random.IntSliceBetween(0, 99999)[0], // Assume this generates a single random number between 0 and 99999
		5,
		4,
		1,
		0,
	}

	fmt.Printf("%sRunning FindPrevPrime tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the FindPrevPrime function in solution/main.go to get the expected result
			expectedResult := solution.FindPrevPrime(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%d'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %d%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the FindPrevPrime function in main.go
			mainResult := FindPrevPrime(val)
			fmt.Printf("%sYour Output: %d%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%d'%s\n\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour FindPrevPrime() = %d\nExpected = %d", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
