package main

import (
	"checkpoint/iscapitalized/solution" // Replace with the correct package path
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

func TestIsCapitalized(t *testing.T) {
	// Predefined test cases
	args := []string{
		"Hello! €How are you?",
		"a",
		"z",
		"!",
		"Hello How Are 4you",
		"What's this 4?",
		"Whatsthis4",
		"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890",
		"131!",
		"H3110 W0r1d!",
		"",
		" ",
	}

	fmt.Printf("%sRunning IsCapitalized tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the IsCapitalized function in solution/main.go to get the expected result
			expectedResult := solution.IsCapitalized(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %t%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the IsCapitalized function in main.go
			mainResult := IsCapitalized(val)
			fmt.Printf("%sYour Output: %t%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour IsCapitalized() = %t\nExpected = %t", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
