package main

import (
    "bytes"
    "os/exec"
    "testing"
)

func TestOutput(t *testing.T) {
    // Execute the Go program directly using 'go run'
    cmd := exec.Command("go", "run", "main.go") // Replace 'main.go' with the appropriate file path if different
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        t.Fatalf("Failed to run the program: %v", err)
    }

    expected := "1\n"
    if out.String() != expected {
        t.Errorf("Unexpected output: got %q, want %q", out.String(), expected)
    }
}