package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"checkpoint/printmemory/solution" // Replace with the correct package path
	"fmt"
	"strconv"

	"github.com/01-edu/go-tests/lib/random" // Assuming the random package is in the same directory or accessible in your GOPATH
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

func TestPrintMemory(t *testing.T) {
	fmt.Printf("%sRunning PrintMemory tests with random and predefined inputs:%s\n", ColorCyan, ColorReset)

	// Array to hold all test cases
	var testCases [][10]byte

	// Generate random test cases
	for j := 0; j < 5; j++ {
		var table [10]byte
		for i := 0; i < random.IntBetween(7, 10); i++ {
			table[i] = byte(random.IntBetween(13, 126))
		}
		testCases = append(testCases, table)
	}

	// Predefined test case
	testCases = append(testCases, [10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	testCases = append(testCases, [10]byte{' ', ' ', ' ', ' ', ' ', 16, 21, '*'})

	// Iterate over all test cases
	for i, table := range testCases {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Run the PrintMemory function in solution/main.go to get the expected result
			expectedResult := captureOutput(func() {
				solution.PrintMemory(table)
			})

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = %v%s\n", ColorBlue, i+1, table, ColorReset)
			fmt.Printf("%sExpected Output:\n%s%s\n", ColorYellow, expectedResult, ColorReset)

			// Run the PrintMemory function in main.go
			mainResult := captureOutput(func() {
				PrintMemory(table)
			})
			fmt.Printf("%sYour Output:\n%s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input %v%s\n\n", ColorRed, ColorBold, table, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour PrintMemory() output =\n%s\nExpected =\n%s", ColorRed, ColorReset, mainResult, expectedResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}

// Helper function to capture the output of a function using os.Pipe
func captureOutput(f func()) string {
	r, w, _ := os.Pipe()
	stdout := os.Stdout
	os.Stdout = w

	f()

	w.Close()
	var buf bytes.Buffer
	io.Copy(&buf, r)
	os.Stdout = stdout

	return buf.String()
}
