package main

import (
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		os.Stdout.WriteString("\n")
		return
	}
	os.Stdout.WriteString(strings.Join(strings.Fields(os.Args[1]), " ") + "\n")

}
