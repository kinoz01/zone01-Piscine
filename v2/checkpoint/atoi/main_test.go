package main

import (
	"checkpoint/atoi/solution"
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
	ColorBlue   = "\033[34m"
	ColorBold   = "\033[1m"
	ColorUnderline = "\033[4m"
)

// Initialize the table with random integers and special test cases
func generateTable() []string {
	table := make([]string, 30)
	for i := range table {
		table[i] = strconv.Itoa(random.Int())
	}
	table = append(table,
		strconv.Itoa(random.MinInt),
		strconv.Itoa(random.MaxInt),
		"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"+657",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
		"123a45",
		"+1234",
		"-1234",
		"+123456",
		"-123456",
	)
	return table
}

func TestAtoiFunctions(t *testing.T) {
	table := generateTable()

	fmt.Printf("%sRunning Atoi tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range table {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			fmt.Printf("%sTest %d: Input = '%s'%s\n", ColorBlue, i+1, val, ColorReset)

			// Run the Atoi function in main.go
			mainResult := Atoi(val)
			fmt.Printf("%sYour Atoi Output: %d%s\n", ColorYellow, mainResult, ColorReset)

			// Run the Atoi function in solution/main.go
			solutionResult := solution.Atoi(val)
			fmt.Printf("%sSolution Atoi Output: %d%s\n", ColorYellow, solutionResult, ColorReset)

			// Compare the results
			if mainResult != solutionResult {
				fmt.Printf("%s%sMismatch for input '%s'%s\n", ColorRed, ColorBold, val, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour Atoi() = %d\nSolution Atoi() = %d", ColorRed, ColorReset, mainResult, solutionResult)
			} else {
				fmt.Printf("%s%sMatch! The outputs are identical.%s\n\n", ColorGreen, ColorBold, ColorReset)
			}
		})
	}
}