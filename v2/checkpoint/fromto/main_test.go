package main

import (
	"checkpoint/fromto/solution" // Replace with the correct package path
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

func TestFromTo(t *testing.T) {
	// Predefined test cases
	args := [][]int{
		{1, 2},
		{1, 10},
		{1, 100},
		{100, 1},
		{99, 1},
		{1, 1},
		{1, 0},
		{0, 1},
		{-1, 10},
		{1, 99},
		{99, -1},
	}

	fmt.Printf("%sRunning FromTo tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, pair := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Extract from and to values
			from, to := pair[0], pair[1]

			// Run the FromTo function in solution/main.go to get the expected result
			expectedResult := solution.FromTo(from, to)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {from: %d, to: %d}%s\n", ColorBlue, i+1, from, to, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the FromTo function in main.go
			mainResult := FromTo(from, to)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {from: %d, to: %d}%s\n\n", ColorRed, ColorBold, from, to, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour FromTo() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
