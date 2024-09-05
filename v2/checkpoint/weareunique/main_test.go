package main

import (
	"checkpoint/weareunique/solution" // Replace with the correct package path
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

func TestWeAreUnique(t *testing.T) {
	// Predefined test cases
	args := [][]string{
		{"abc", "def"},
		{"hello", "yoall"},
		{"everyone", ""},
		{"hello world", "fam"},
		{"abc", "abc"},
		{"", ""},
		{"pomme", "pomme"},
		{"+265", "265"},
		{"123231", "123231"},
		{"w^p@@j", "w^p@@j"},
		{"26235e5", "4478q92"},
		{"		", "		 "},
		{"AB$%d.52", "eepqdl.52"},
		{"", "eveRyone"},
		{"_55w1se", "55w1se"},
	}

	fmt.Printf("%sRunning WeAreUnique tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Extract str1 and str2
			str1, str2 := val[0], val[1]

			// Run the WeAreUnique function in solution/main.go to get the expected result
			expectedResult := solution.WeAreUnique(str1, str2)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {str1: '%s', str2: '%s'}%s\n", ColorBlue, i+1, str1, str2, ColorReset)
			fmt.Printf("%sExpected Output: %d%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the WeAreUnique function in main.go
			mainResult := WeAreUnique(str1, str2)
			fmt.Printf("%sYour Output: %d%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {str1: '%s', str2: '%s'}%s\n\n", ColorRed, ColorBold, str1, str2, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour WeAreUnique() = %d\nExpected = %d", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
