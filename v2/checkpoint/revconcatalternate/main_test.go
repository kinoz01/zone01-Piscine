package main

import (
	"checkpoint/revconcatalternate/solution" // Replace with the correct package path
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

func TestRevConcatAlternate(t *testing.T) {
	// Predefined test cases
	testCases := []struct {
		args [][]int
		want []int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[]int{3, 6, 2, 5, 1, 4},
		},
		{
			[][]int{{4, 5, 6, 7, 8, 9}, {1, 2, 3}},
			[]int{9, 8, 7, 6, 3, 5, 2, 4, 1},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6, 7, 8, 9}},
			[]int{9, 8, 7, 3, 6, 2, 5, 1, 4},
		},
		{
			[][]int{{1, 2, 3}, {4, 5}},
			[]int{3, 2, 5, 1, 4},
		},
		{
			[][]int{{}, {4, 5, 6}},
			[]int{6, 5, 4},
		},
		{
			[][]int{{1, 2, 3}, {}},
			[]int{3, 2, 1},
		},
		{
			[][]int{{}, {}},
			[]int{},
		},
		{
			[][]int{{1, 2, 4}, {10, 20, 30, 40, 50}},
			[]int{50, 40, 4, 30, 2, 20, 1, 10},
		},
	}

	fmt.Printf("%sRunning RevConcatAlternate tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, tc := range testCases {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the RevConcatAlternate function in solution/main.go to get the expected result
			expectedResult := solution.RevConcatAlternate(tc.args[0], tc.args[1])

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {slice1: %v, slice2: %v}%s\n", ColorBlue, i+1, tc.args[0], tc.args[1], ColorReset)
			fmt.Printf("%sExpected Output: %v%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the RevConcatAlternate function in main.go
			mainResult := RevConcatAlternate(tc.args[0], tc.args[1])
			fmt.Printf("%sYour Output: %v%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if fmt.Sprintf("%v", mainResult) != fmt.Sprintf("%v", expectedResult) {
				fmt.Printf("%s%sMismatch for input {slice1: %v, slice2: %v}%s\n", ColorRed, ColorBold, tc.args[0], tc.args[1], ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour RevConcatAlternate() = %v\nExpected = %v", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
