package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		FileContent, error := ioutil.ReadFile(os.Args[1])
		if error != nil {
			fmt.Println("Wrong Filename")
		}
		fmt.Printf("%s", FileContent)
	}
}

// quest 8
