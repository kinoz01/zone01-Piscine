package main

import (
	"bytes"
	"checkpoint/chunk/solution" // Replace with the correct package path
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"

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
)

// Function to generate a random size slice of integers
func randomSize() []int {
	var randSlice []int
	for i := 0; i <= random.IntBetween(0, 20); i++ {
		randSlice = append(randSlice, random.Int())
	}
	return randSlice
}

func TestChunk(t *testing.T) {
	// Predefined test cases
	type node struct {
		slice []int
		ch    int
	}
	args := []node{
		{
			slice: []int{},
			ch:    0,
		},
		{
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8},
			ch:    0,
		},
	}

	// Adding random test cases
	for i := 0; i <= 7; i++ {
		value := node{
			slice: randomSize(),
			ch:    random.IntBetween(0, 10),
		}
		args = append(args, value)
	}

	fmt.Printf("%sRunning Chunk tests with the following inputs:%s\n", ColorCyan, ColorReset)

	for i, val := range args {
		t.Run("Test_"+strconv.Itoa(i), func(t *testing.T) {
			// Capture the expected result from solution.Chunk
			expectedResult := captureOutput(func() {
				solution.Chunk(val.slice, val.ch)
			})

			// Print the input and the expected result
			fmt.Printf("%sTest %d: Input = {slice: %v, ch: %d}%s\n", ColorBlue, i+1, val.slice, val.ch, ColorReset)
			fmt.Printf("%sExpected Output:\n%s%s\n", ColorYellow, expectedResult, ColorReset)

			// Capture the result from the main Chunk function
			mainResult := captureOutput(func() {
				Chunk(val.slice, val.ch)
			})
			fmt.Printf("%sYour Output:\n%s%s\n", ColorYellow, mainResult, ColorReset)

			// Compare the results
			if mainResult != expectedResult {
				fmt.Printf("%s%sMismatch for input {slice: %v, ch: %d}%s\n", ColorRed, ColorBold, val.slice, val.ch, ColorReset)
				t.Errorf("%sMismatch Details:%s\nYour Chunk() output =\n%s\nExpected =\n%s", ColorRed, ColorReset, mainResult, expectedResult)
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
