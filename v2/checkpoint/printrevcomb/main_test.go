package main

import (
	"bytes"
	"os/exec"
	"testing"
)

func runProgram(path string) (string, error) {
	cmd := exec.Command("go", "run", path)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func TestProgramOutput(t *testing.T) {
	// Path to the main.go in the solution folder
	solutionPath := "./solution/main.go" // Adjust the path as necessary
	// Path to the main.go outside the solution folder
	externalPath := "./main.go" // Adjust the path as necessary

	// Run the program in the solution folder
	solutionOutput, err := runProgram(solutionPath)
	if err != nil {
		t.Fatalf("Failed to run solution program: %v", err)
	}

	// Run the external program
	externalOutput, err := runProgram(externalPath)
	if err != nil {
		t.Fatalf("Failed to run external program: %v", err)
	}

	// Compare the outputs
	if solutionOutput != externalOutput {
		t.Errorf("Output mismatch:\nExpected:\n%s\nGot:\n%s", externalOutput, solutionOutput)
	}
}
