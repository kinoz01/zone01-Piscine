package main

import (
	"fmt"
	"strconv"
	"testing"

	"checkpoint/itoa/solution"

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

func TestItoa(t *testing.T) {
	fmt.Printf("%sRunning Itoa tests with random inputs:%s\n", ColorCyan, ColorReset)

	for i := 0; i < 50; i++ {
		val := random.Int() // Generate a random integer

		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the Itoa function in solution/main.go to get the expected result
			expectedResult := solution.Itoa(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = %d%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the Itoa function in main.go
			mainResult := Itoa(val)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input %d%s\n\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour Itoa() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
