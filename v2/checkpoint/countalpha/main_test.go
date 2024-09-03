package main

import (
	"checkpoint/countalpha/solution"
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

func TestExampleFunction(t *testing.T) {
	// Predefined test cases
	params := []string{
		"a",
		"abc",
		"",
		"H E L L O",
		"A1B2C3D4B5DE6",
		"a1b2c3d4b5de6",
		"A1b2C3D4B5De6",
		"1234",
		"    ",
	}

	fmt.Printf("%sRunning ExampleFunction tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range params {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the ExampleFunction function in solution/main.go to get the expected result
			expectedResult := solution.CountAlpha(val) // Replace ExampleFunction with the actual function name

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %v%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the ExampleFunction function in main.go
			mainResult := CountAlpha(val) // Replace ExampleFunction with your actual function name
			fmt.Printf("%sYour Output: %v%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour ExampleFunction() = %v\nExpected = %v", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
