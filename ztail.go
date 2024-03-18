package main

import (
	"fmt"
	"os"
)

func isnum(s string) bool {
	for _, digit := range s {
		if digit > '9' || digit < '0' {
			return false
		}
	}
	return true
}

func atoii(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		os.Exit(1)
	} else if args[0] != "-c" {
		os.Exit(1)
	} else if !isnum(args[1]) {
		os.Exit(1)
	}
	x := atoii(args[1])
	errorOccurred := false // this is added so we do continue our loop and we gonna use it lastly to exit lastely using os.Exit(1)
	for i := 2; i < len(args); i++ { // this loop is for checking errors before printing
		filePath := args[i]
		_, err := os.Open(filePath)
		if err != nil {
			errorOccurred = true
		}
	}
	for i := 2; i < len(args); i++ {
		filePath := args[i]
		file, err := os.Open(filePath) // open file
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", args[i])
			if errorOccurred && len(args) >= 4 && i != len(args)-1 {
				fmt.Printf("\n")
			}
		} else {
			defer file.Close() // close file
			fileInfo, err := file.Stat() // get file stat
			if err != nil {
				errorOccurred = true
			}
			fileSize := fileInfo.Size() // get file size
			fileContent := make([]byte, fileSize)
			_, err = file.Read(fileContent) // read file and put it's content in fileContent
			if err != nil {
				errorOccurred = true
			}
			runeFile := []rune(string(fileContent))
			if x > len(runeFile) {
				fmt.Printf("Number %v exceed %s size\n", x, args[i])
				errorOccurred = true
				if i != len(args)-1 {
					fmt.Printf("\n")
				}
			} else {
				if len(args) != 3 {
					fmt.Printf("==> %s <==\n", args[i])
				}
				for j := len(runeFile) - x; j < len(runeFile); j++ {
					fmt.Printf("%c", runeFile[j])
				}
				if i != len(args)-1 && !errorOccurred {
					fmt.Printf("\n")
				}
			}
		}
	}
	if errorOccurred {
		os.Exit(1)
	}
}

// quest 8

/* Note: the code is simple but a lot of "if"s are
added just to keep up with the wanted printed results in
the examples provided. 
Restriction on using os package and not using ioutil 
package like we did in cat.go also made the code longer. */
