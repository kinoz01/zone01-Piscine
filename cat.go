package main

import (
	"io"
	"io/ioutil"
	"os"
	"github.com/01-edu/z01"
)

func printstrs(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	}
	for i := 0; i < len(args); i++ {
		filecontent, err := ioutil.ReadFile(args[i])
		if err != nil {
			printstrs("ERROR: open ")
			printstrs(args[i])
			printstrs(": No such file or directory\nexit status 1\n")
		} else {
			printstrs(string(filecontent)) // filecontent is a slice of byte so we transform it to a string
		}
	}
}

// quest 8
