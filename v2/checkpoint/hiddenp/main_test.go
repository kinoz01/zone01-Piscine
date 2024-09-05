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
)

// Initial test cases based on the provided string pairs
var testCases = []struct {
	args [2]string
}{
	{[2]string{"fgex.;", "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"}},
	{[2]string{"abc", "2altrb53c.sse"}},
	{[2]string{"abc", "btarc"}},
	{[2]string{"DD", "DABC"}},
	{[2]string{""}},
}

// Add random string pairs to the test cases
func init() {
	for i := 0; i < 5; i++ {
		testCases = append(testCases,
			struct{ args [2]string }{[2]string{random.Str(chars.Lower, 1), random.Str(chars.Lower, 13)}},
			struct{ args [2]string }{[2]string{random.Str(chars.Upper, 1), random.Str(chars.Upper, 13)}},
		)
	}
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
func printCaseDetails(caseNumber int, args [2]string, expectedOutput, yourOutput string) {
	fmt.Printf("%sCase Number:%s %d\n", ColorBlue, ColorReset, caseNumber)
	fmt.Printf("%sInput:%s %v\n", ColorYellow, ColorReset, args)
	fmt.Printf("%sExpected Output:%s %s", ColorGreen, ColorReset, expectedOutput)
	fmt.Printf("%sYour Output:%s %s\n", ColorBlue, ColorReset, yourOutput)
}

func TestPrograms(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1), func(t *testing.T) {
			// Run the first program located in the "solution" directory to get expected output
			expectedOutput, err := runGoFile("solution", "main.go", tc.args[:]...)
			if err != nil {
				t.Fatalf("%sSolution program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Run the second program located in the current directory
			mainOutput, err := runGoFile(".", "main.go", tc.args[:]...)
			if err != nil {
				t.Fatalf("%sMain program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Print the case details with both expected and your output
			printCaseDetails(i+1, tc.args, expectedOutput, mainOutput)

			// Compare the outputs
			if mainOutput != expectedOutput {
				t.Errorf("%sOutputs differ for input %v:\n%sExpected (Solution):\n%s%s\n%sGot (Main):\n%s%s",
					ColorRed, tc.args, ColorYellow, ColorReset, expectedOutput, ColorYellow, ColorReset, mainOutput)
			} else {
				fmt.Printf("%sTest Case %d passed!%s\n\n", ColorGreen, i+1, ColorReset)
			}
		})
	}
}
