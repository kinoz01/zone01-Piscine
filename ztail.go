package main

import (
	"fmt"
	"os"
)

var args []string

func init() {
	args = os.Args[1:]
}

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

func exitStatus2(flag string) {
	if flag == "-c" && len(args) == 1 {
		fmt.Printf("flag needs an argument: %s\n", flag)
	} else if flag != "-c" && len(args) == 1 {
		fmt.Printf("flag provided but not defined: %s\n", flag)
	} else if len(args) >= 2 {
		fmt.Printf("invalid value \"%s\" for flag -c: parse error\n", flag)
	}
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("  -c int\n\toutput the last NUM bytes\n")
	os.Exit(2)
}

func printnl(i int, errorindex []int) {
	if i != len(args)-1 {
		c := 0
		for k := 0; k < len(errorindex); k++ {
			if i == errorindex[k] {
				c = 1
				break
			}
		}
		if c == 0 {
			fmt.Printf("\n")
		}
	}
	if len(args) == 3 {
		fmt.Printf("\n")
	}
}

func main() {
	/**************** Handling Invalid Input ****************/
	if len(args) == 0 {
		return
	}
	p1 := []rune(args[0])
	if len(p1) == 0 || p1[0] != '-' {
		for i := 0; i < len(args); i++ {
			fmt.Printf("open %s: no such file or directory\n", args[i])
		}
		os.Exit(1)
	}
	if len(args) == 1 {
		exitStatus2(string(p1))
	}
	if !isnum(args[1]) {
		exitStatus2(args[1])
	}
	/*******************************************************/
	x := atoii(args[1])
	errorOccurred := false // this is added so we do continue our loop and we gonna use it to exit lastely using os.Exit(1)
	errorindex := []int{}
	for i := 2; i < len(args); i++ { // this loop is for checking errors before printing
		filePath := args[i]
		_, err := os.Open(filePath)
		if err != nil {
			errorOccurred = true
			errorindex = append(errorindex, i-1)
		}
	}
	// fmt.Println(errorindex)
	for i := 2; i < len(args); i++ {
		filePath := args[i]
		file, err := os.Open(filePath) // open file
		if err != nil {
			fmt.Printf("open %s: no such file or directory\n", args[i])
			printnl(i, errorindex)
		} else {
			defer file.Close()           // close file
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

			if len(args) != 3 {
				fmt.Printf("==> %s <==\n", args[i])
			}
			if len(runeFile) <= x {
				fmt.Printf(string(fileContent))
				printnl(i, errorindex)
			} else {
				for j := len(runeFile) - x; j < len(runeFile); j++ {
					fmt.Printf("%c", runeFile[j])
				}
				printnl(i, errorindex)
			}
		}
	}
	if errorOccurred {
		os.Exit(1)
	}
}

// quest 8

/* Note: the core of the code is simple but a lot of
code is added just to keep up with the wanted
printed results in the examples provided.
Restriction on using os package and not using ioutil
package like we did in cat.go also made the code longer.
*/
