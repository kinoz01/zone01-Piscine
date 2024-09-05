package main

import (
	"checkpoint/fifthandskip/solution" // Replace with the correct package path
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

func TestFifthAndSkip(t *testing.T) {
	// Predefined test cases
	table := []string{
		"1234556789", 
		"e 5£ @ 8* 7 =56 ;", 
		"QKplq%QSw", 
		"", 
		"hello \\! n4ght cr3a8ure7 ", 
		"Kimetsu no Yaiba", 
		"8595485-52", 
		"-552", 
		"w58tw7474abc", 
		"Po65 4o",
	}

	fmt.Printf("%sRunning FifthAndSkip tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range table {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the FifthAndSkip function in solution/main.go to get the expected result
			expectedResult := solution.FifthAndSkip(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the FifthAndSkip function in main.go
			mainResult := FifthAndSkip(val)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour FifthAndSkip() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
