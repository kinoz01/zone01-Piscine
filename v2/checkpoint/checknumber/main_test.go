package main

import (
	"checkpoint/checknumber/solution"
	"fmt"
	"strconv"
	"testing"
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

func TestCheckNumberFunctions(t *testing.T) {
	// Predefined test cases
	args := []string{
		"123",
		"H1ll0",
		"",
		"1",
		"1.1",
		"Containe1number",
		"     ",
		"upson lorem ipsum",
	}

	fmt.Printf("%sRunning CheckNumber tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the CheckNumber function in solution/main.go to get the expected result
			expectedResult := solution.CheckNumber(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected CheckNumber Output: %t%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the CheckNumber function in main.go
			mainResult := CheckNumber(val)
			fmt.Printf("%sYour CheckNumber Output: %t%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour CheckNumber() = %t\nExpected CheckNumber() = %t", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
