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

// Test cases based on your provided examples and additional multi-argument cases
var testCases = []struct {
	args []string
}{
	{[]string{"(johndoe)"}},
	{[]string{")()"}},
	{[]string{"([)]"}},
	{[]string{"{2*[d - 3]/(12)}"}},
	{[]string{"", "{[(0 + 0)(1 + 1)](3*(-1)){()}}"}},
	{[]string{"{][]}", "{3*[21/(12+ 23)]}"}},
	{[]string{"{([)])}", "{{{something }- [something]}}", "there are"}},
}

// Generate 18 random tests (at least half are valid)
func init() {
	for i := 0; i < 3; i++ {
		testCases = append(testCases,
			struct{ args []string }{[]string{"(" + random.Str(chars.ASCII, 13) + ")"}},
			struct{ args []string }{[]string{"[" + random.Str(chars.ASCII, 13) + "]"}},
			struct{ args []string }{[]string{"{" + random.Str(chars.ASCII, 13) + "}"}},
			struct{ args []string }{[]string{"(" + random.Str(chars.Alnum, 13) + ")"}},
			struct{ args []string }{[]string{"[" + random.Str(chars.Alnum, 13) + "]"}},
			struct{ args []string }{[]string{"{" + random.Str(chars.Alnum, 13) + "}"}},
		)
	}

	testCases = append(testCases,
		struct{ args []string }{[]string{"(" + random.Str(chars.ASCII, 13) + ")"}},
		struct{ args []string }{[]string{"[" + random.Str(chars.ASCII, 13) + "]"}},
		struct{ args []string }{[]string{"{" + random.Str(chars.ASCII, 13) + "}"}},
		struct{ args []string }{[]string{"(" + random.Str(chars.Alnum, 13) + ")"}},
		struct{ args []string }{[]string{"[" + random.Str(chars.Alnum, 13) + "]"}},
		struct{ args []string }{[]string{"{" + random.Str(chars.Alnum, 13) + "}"}},
	)
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
func printCaseDetails(caseNumber int, args []string, expectedOutput string) {
	fmt.Printf("%sCase Number:%s %d\n", ColorBlue, ColorReset, caseNumber)
	fmt.Printf("%sInput:%s %v\n", ColorYellow, ColorReset, args)
	fmt.Printf("%sExpected Output:%s %s", ColorGreen, ColorReset, expectedOutput)
}

func TestPrograms(t *testing.T) {
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case %d", i+1), func(t *testing.T) {
			// Run the first program located in the "solution" directory to get expected output
			expectedOutput, err := runGoFile("solution", "main.go", tc.args...)
			if err != nil {
				t.Fatalf("%sSolution program failed:%s %v", ColorRed, ColorReset, err)
			}

			// Print the case details
			printCaseDetails(i+1, tc.args, expectedOutput)

			// Run the second program located in the current directory
			mainOutput, err := runGoFile(".", "main.go", tc.args...)
			if err != nil {
				t.Fatalf("%sMain program failed:%s %v", ColorRed, ColorReset, err)
			}

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