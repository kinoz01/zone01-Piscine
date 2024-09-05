package main

import (
	"checkpoint/thirdtimeisacharm/solution" // Replace with the correct package path
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

func TestThirdTimeIsACharm(t *testing.T) {
	// Predefined test cases
	args := []string{
		"1234556789",
		"QKplq%QSw",
		"",
		"Kimetsu no Yaiba",
		"Z",
		"email123@live.fr",
		"8595485-52",
		"-552",
		"abc",
		"w58tw7474abc",
		"fifa world cup `2022`",
	}

	fmt.Printf("%sRunning ThirdTimeIsACharm tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the ThirdTimeIsACharm function in solution/main.go to get the expected result
			expectedResult := solution.ThirdTimeIsACharm(val)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the ThirdTimeIsACharm function in main.go
			mainResult := ThirdTimeIsACharm(val)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour ThirdTimeIsACharm() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
