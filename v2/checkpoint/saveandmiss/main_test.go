package main

import (
	"checkpoint/saveandmiss/solution" // Replace with the correct package path
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

func TestSaveAndMiss(t *testing.T) {
	// Predefined test cases
	testCases := []struct {
		arg string
		num int
	}{
		{"123456789", 1},
		{"e 5£ @ 8* 7 =56 ;", 2},
		{"QKplq%QSw", 3},
		{"", 4},
		{"hello \\! n4ght cr3a8ure7 ", 5},
		{"Kimetsu no Yaiba", 6},
		{"8595485-52", 7},
		{"abcdefghijklmnopqrstuvwyz", 8},
		{"w58tw7474abc", 9},
		{"Po65 4o", 10},
	}

	fmt.Printf("%sRunning SaveAndMiss tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, tc := range testCases {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the SaveAndMiss function in solution/main.go to get the expected result
			expectedResult := solution.SaveAndMiss(tc.arg, tc.num)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {arg: '%s', num: %d}%s\n", ColorBlue, i+1, tc.arg, tc.num, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the SaveAndMiss function in main.go
			mainResult := SaveAndMiss(tc.arg, tc.num)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {arg: '%s', num: %d}%s\n", ColorRed, ColorBold, tc.arg, tc.num, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour SaveAndMiss() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
