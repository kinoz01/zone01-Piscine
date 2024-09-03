package main

import (
	"checkpoint/cameltosnakecase/solution" // Replace with the correct package path
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

func TestCamelToSnakeCase(t *testing.T) {
	// Predefined test cases
	args := []string{
		"CamelCase",
		"camelCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCEf",
		"abcAree",
		"ahe1Abde",
		"tesTing1",
		"SOME_VARIABLE",
		"ASuperLonGVariableName",
		"thisIsaTestOfCamelCase",
		"aA",
	}

	fmt.Printf("%sRunning CamelToSnakeCase tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the CamelToSnakeCase function in solution/main.go to get the expected result
			expectedResult := solution.CamelToSnakeCase(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the CamelToSnakeCase function in main.go
			mainResult := CamelToSnakeCase(val)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour CamelToSnakeCase() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
