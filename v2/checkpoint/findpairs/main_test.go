package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"
)

// ANSI escape codes for coloring
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
)

// Test cases based on your provided struct
var testCases = []struct {
	arr       string
	targetSum string
}{
	{arr: "[1, 2, 3, 4, 5]", targetSum: "6"},
	{arr: "[1, 2, 3, 4, 5, 1]", targetSum: "6"},
	{arr: "[-1, 2, -3, 4, -5]", targetSum: "1"},
	{arr: "[-1, -2, -3, -4, -5]", targetSum: "-5"},
	{arr: "[1, 2, 3, 4, 5]", targetSum: "10"},
	{arr: "[1, 2, 3, 4, 20, -4, 5]", targetSum: "2 5"},
	{arr: "[1, 2, 3, 4, 20, -4, 5]", targetSum: "l"},
	{arr: "[1, 2, 3, 4, 20, p, 5]", targetSum: "5"},
	{arr: "[1, 2, 3, 4, 20,    p, 5]", targetSum: "5"},
	{arr: "[1, 2, 3, 4, 20, 5", targetSum: "5"},
	{arr: "1, 2, 3, 4, 20, 5", targetSum: "5"},
	{arr: "1 2 3 4 20 5", targetSum: "5"},
}

// Helper function to run a Go file as a separate process and capture the output
func runGoFile(dir, filename string, args ...string) (string, error) {
	cmd := exec.Command("go", append([]string{"run", filename}, args...)...)
	cmd.Dir = dir
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

// Function to print the case number, input, and expected output
func printCaseDetails(caseNumber int, arr, targetSum, expectedOutput, yourOutput string) {
	fmt.Printf("%sCase Number:%s %d\n", ColorBlue, ColorReset, caseNumber)
	fmt.Printf("%sArray Input:%s %s\n", ColorYellow, ColorReset, arr)
	fmt.Printf("%sTarget Sum:%s %s\n", ColorYellow, ColorReset, targetSum)
	fmt.Printf("%sExpected Output:%s %s", ColorGreen, ColorReset, expectedOutput)
	fmt.Printf("%sYour Output:%s %s\n", ColorBlue, ColorReset, yourOutput)
}

func TestPrograms(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1), func(t *testing.T) {
			// Run the first program located in the "solution" directory to get expected output
			expectedOutput, err := runGoFile("solution", "main.go", tc.arr, tc.targetSum)
			if err != nil {
				t.Fatalf("%sSolution program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Run the second program located in the current directory
			mainOutput, err := runGoFile(".", "main.go", tc.arr, tc.targetSum)
			if err != nil {
				t.Fatalf("%sMain program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Print the case details with both expected and your output
			printCaseDetails(i+1, tc.arr, tc.targetSum, expectedOutput, mainOutput)

			// Compare the outputs
			if mainOutput != expectedOutput {
				t.Errorf("%sOutputs differ for input arr: %v targetSum: %v\n%sExpected (Solution):\n%s%s\n%sGot (Main):\n%s%s",
					ColorRed, tc.arr, tc.targetSum, ColorYellow, ColorReset, expectedOutput, ColorYellow, ColorReset, mainOutput)
			} else {
				fmt.Printf("%sTest Case %d passed!%s\n\n", ColorGreen, i+1, ColorReset)
			}
		})
	}
}
