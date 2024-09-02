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
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

// Test cases for the Reverse Polish Notation (RPN) calculator
var testCases = []struct {
	input    string
	expected string
}{
	{"3 4 +", "7\n"},
	{"10 2 /", "5\n"},
	{"2 3 + 5 *", "25\n"},
	{"5 1 2 + 4 * + 3 -", "14\n"},
	{"15 7 1 1 + - / 3 * 2 1 1 + + -", "5\n"},
	{"5 1 2 + 4 * + 3 - 7 /", "Error\n"},
	{"1 2 3 +", "Error\n"},
	{"3 4", "Error\n"},
	{"3 +", "Error\n"},
	{"", "Error\n"},
	{"3 3 + +", "Error\n"},
	{"100 200 + 2 / 5 * 7 +", "757\n"},
	{"5 2 4 * + 7 -", "Error\n"},
}

// Helper function to run the Go program and capture the output
func runGoProgram(args string) (string, error) {
	cmd := exec.Command("go", "run", "main.go", args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func TestRPNCalculator(t *testing.T) {
	fmt.Printf("%sRunning tests with the following cases:%s\n", ColorCyan, ColorReset)
	for i, tc := range testCases {
		fmt.Printf("%sTest Case %d%s:\n", ColorBlue, i+1, ColorReset)
		fmt.Printf("%sInput = '%s', Expected Output = '%s'%s\n", ColorYellow, tc.input, tc.expected, ColorReset)
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1), func(t *testing.T) {
			output, err := runGoProgram(tc.input)
			if err != nil {
				t.Fatalf("%sProgram execution failed:%s %v", ColorRed, ColorReset, err)
			}

			// Compare the outputs
			if output != tc.expected {
				t.Errorf("%sOutputs differ:\n%sExpected:\n%s%s\n%sGot:\n%s%s", ColorRed, ColorYellow, ColorReset, tc.expected, ColorYellow, ColorReset, output)
			} else {
				fmt.Printf("%sOutputs match for Test Case %d!%s\n", ColorGreen, i+1, ColorReset)
			}
		})
	}
}
