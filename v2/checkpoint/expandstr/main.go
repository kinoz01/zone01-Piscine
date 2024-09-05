package main

import (
	"os"
	"strings"
)

func Fields(s string) []string {
	temp := []rune{}
	result := []string{}
	for _, char := range s {
		if char != ' ' {
			temp = append(temp, char)
		} else if len(temp) > 0 {
			result = append(result, string(temp))
			temp = nil
		}
	}
	if len(temp) > 0 {
		result = append(result, string(temp))
	}
	return result
}

func main() {
	if len(os.Args) != 2 || os.Args[1] == "" {
		return
	}
	slic := Fields(os.Args[1])
	if len(slic) == 0 {
		return
	}
	result := slic[0]
	for i := 1; i < len(slic); i++ {
		result += "   " + slic[i]
	}
	os.Stdout.WriteString(result + "\n")
}

// Solution Algo (Run it in main to test it)
func IfStringsAllowed() {
	if len(os.Args) != 2 {
		return
	}
	os.Stdout.WriteString(strings.Join(strings.Fields(os.Args[1]), "   ") + "\n")
}
