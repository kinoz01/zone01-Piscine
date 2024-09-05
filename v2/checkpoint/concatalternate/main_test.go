package main

import (
	"checkpoint/concatalternate/solution" // Replace with the correct package path
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

func TestConcatAlternate(t *testing.T) {
	// Predefined test cases
	testCases := []struct {
		slice1 []int
		slice2 []int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{1, 2, 3}, []int{4, 5}},
		{[]int{}, []int{4, 5, 6}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{}, []int{}},
		{[]int{1, 2, 4}, []int{10, 20, 30, 40, 50}},
	}

	fmt.Printf("%sRunning ConcatAlternate tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, tc := range testCases {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the ConcatAlternate function in solution/main.go to get the expected result
			expectedResult := solution.ConcatAlternate(tc.slice1, tc.slice2)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {slice1: %v, slice2: %v}%s\n", ColorBlue, i+1, tc.slice1, tc.slice2, ColorReset)
			fmt.Printf("%sExpected Output: %v%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the ConcatAlternate function in main.go
			mainResult := ConcatAlternate(tc.slice1, tc.slice2)
			fmt.Printf("%sYour Output: %v%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if fmt.Sprintf("%v", mainResult) != fmt.Sprintf("%v", expectedResult) {
				fmt.Printf("%s%sMismatch for input {slice1: %v, slice2: %v}%s\n", ColorRed, ColorBold, tc.slice1, tc.slice2, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour ConcatAlternate() = %v\nExpected = %v", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
