package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"testing"

	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
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

// Test cases as per your input examples
var testCases = []struct {
	regex string
	text  string
}{
	{"(a)", "I'm heavyjumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"},
	{"(e|n)", "I currently have 4 windows opened up… and I don’t know why."},
	{"(hi)", "He swore he just saw his sushi move."},
	{"(s)", ""},
	{"i", "Something in the air"},
	{"(|)", "hello"},
	{"(|)", ""},
	{validRegExp(2), random.Str(chars.Words, 60)},
	{validRegExp(2), random.Str(chars.Words, 60)},
	{validRegExp(6), random.Str(chars.Words, 60)},
	{random.Str("axyz", 1), random.Str("axyzdassbzzxxxyy cdq     ", 10)},
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

func TestPrograms(t *testing.T) {
	fmt.Printf("%sRunning tests with the following cases:%s\n", ColorCyan, ColorReset)
	for i, tc := range testCases {
		fmt.Printf("%sTest Case %d%s:\n", ColorBlue, i+1, ColorReset)
		fmt.Printf("%sRegex = '%s', %s\n", ColorYellow, tc.regex, ColorReset)
		fmt.Printf("%sText = '%s'%s\n", ColorYellow, tc.text, ColorReset)
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1,), func(t *testing.T) {
			// Run the first program located in the "solution" directory
			firstOutput, err := runGoFile("solution", "main.go", tc.regex, tc.text)
			if err != nil {
				t.Fatalf("%sYour program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Run the second program located in the current directory
			secondOutput, err := runGoFile(".", "main.go", tc.regex, tc.text)
			if err != nil {
				t.Fatalf("%sSecond program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Compare the outputs
			if firstOutput != secondOutput {
				t.Errorf("%sOutputs differ:\n%sYour Program Output:\n%s%s\n%sSolution Program Output:\n%s%s", ColorRed, ColorYellow, ColorReset, firstOutput, ColorYellow, ColorReset, secondOutput)
			} else {
				fmt.Printf("%sOutputs match for Test Case %d!%s\n", ColorGreen, i+1, ColorReset)
			}
		})
	}
}

func validRegExp(n int) string {
	result := "("

	for i := 0; i < n; i++ {
		result += random.Str(chars.Lower, 1)
		if random.Int()%2 == 0 {
			result += random.Str(chars.Lower, 1)
		}
		if i != n-1 {
			result += "|"
		}
	}

	result += ")"
	return result
}