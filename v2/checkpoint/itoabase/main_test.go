package main

import (
	"checkpoint/itoabase/solution" // Replace with the correct package path
	"fmt"
	"strconv"
	"testing"

	"github.com/01-edu/go-tests/lib/random" // Assuming the random package is in the same directory or accessible in your GOPATH
)

// ANSI escape codes for coloring and formatting
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorCyan   = "\033[36m"
	ColorBlue   = "\033/34m"
	ColorBold   = "\033[1m"
)

func TestItoaBase(t *testing.T) {
	// Random test cases
	for i := 0; i < 30; i++ {
		value := random.IntBetween(-1000000, 1000000)
		base := random.IntBetween(2, 16)

		t.Run("Test_Random_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the ItoaBase function in solution/main.go to get the expected result
			expectedResult := solution.ItoaBase(value, base)

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {n: %d, base: %d}%s\n", ColorBlue, i+1, value, base, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the ItoaBase function in main.go
			mainResult := ItoaBase(value, base)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {n: %d, base: %d}%s\n", ColorRed, ColorBold, value, base, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour ItoaBase() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}

	// Special cases for MaxInt and MinInt
	for i := 0; i < 5; i++ {
		base := random.IntBetween(2, 16)

		t.Run("Test_MaxInt_"+strconv.Itoa(i), func(t *testing.T) {
			expectedResult := solution.ItoaBase(random.MaxInt, base)
			fmt.Printf("%sTest MaxInt %d: Input = {n: %d, base: %d}%s\n", ColorBlue, i+1, random.MaxInt, base, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			mainResult := ItoaBase(random.MaxInt, base)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {n: %d, base: %d}%s\n", ColorRed, ColorBold, random.MaxInt, base, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour ItoaBase() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})

		t.Run("Test_MinInt_"+strconv.Itoa(i), func(t *testing.T) {
			expectedResult := solution.ItoaBase(random.MinInt, base)
			fmt.Printf("%sTest MinInt %d: Input = {n: %d, base: %d}%s\n", ColorBlue, i+1, random.MinInt, base, ColorReset)
			fmt.Printf("%sExpected Output: %s%s\n", ColorYellow, expectedResult, ColorReset)

			mainResult := ItoaBase(random.MinInt, base)
			fmt.Printf("%sYour Output: %s%s\n", ColorYellow, mainResult, ColorReset)

			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {n: %d, base: %d}%s\n", ColorRed, ColorBold, random.MinInt, base, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour ItoaBase() = %s\nExpected = %s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}
