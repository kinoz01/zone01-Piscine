package main

import (
	"checkpoint/canjump/solution" // Replace with the correct package path
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

func TestCanJump(t *testing.T) {
	// Predefined test cases
	tests := [][]uint{
		{2, 3, 1, 1, 4},
		{1, 1, 1, 1, 0},
		{5, 4, 3, 2, 1, 0},
		{0},
		{5},
		{},
		{1, 2, 3, 0, 2},
		{3, 2, 1, 0, 4},
		{0, 0, 0, 0, 0},
		{1, 2, 3},
		{1, 2, 3, 0, 1},
		{1, 0, 0, 0, 0},
		{1, 0, 1, 0, 1},
		{10, 20, 30, 40, 0},
	}

	fmt.Printf("%sRunning CanJump tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range tests {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the CanJump function in solution/main.go to get the expected result
			expectedResult := solution.CanJump(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = %v%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %t%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the CanJump function in main.go
			mainResult := CanJump(val)
			fmt.Printf("%sYour Output: %t%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input %v%s\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour CanJump() = %t\nExpected = %t", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
