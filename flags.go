package main

import (
	"os"

	"fmt"
)

func msg(n int) {
	if n == 0 {
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	}
	if n == 1 {
		fmt.Print("--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
	}
	if n == 2 {
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n")
	}
	if n == 3 {
		fmt.Print("--insert\n  -i\n\t Add string to be inserted after = sign.\n")
	}

}

func oorder(s string) string {
	sliceArg := []rune(s)
	for i := 0; i < len(sliceArg); i++ {
		for j := i + 1; j < len(sliceArg); j++ {
			if sliceArg[i] > sliceArg[j] {
				sliceArg[i], sliceArg[j] = sliceArg[j], sliceArg[i]
			}
		}
	}
	return string(sliceArg)
}

func main() {
	args := os.Args[1:]
	order := false
	insert := false
	var insertion string
	var str string
	var str2 string
	if len(args) > 3 {
		return
	}
	if len(args) == 0 || args[0] == "-h" || args[0] == "--help" {
		msg(0)
		return
	}
	for i := 0; i < len(args); i++ {
		if args[i] == "--order" || args[i] == "-o" {
			if args[len(args)-1] == "--order" || args[len(args)-1] == "-o" || args[i+1] == "" {
				msg(1)
				return
			}
			str = args[i+1]
			order = true
		} else if len(args[i]) >= 8 && args[i][:8] == "--insert" {
			if len(args[i]) == 8 || args[i][8] != '=' {
				msg(3)
				return
			}
			insertion = args[i][9:]
			if len(insertion) == 0 {
				msg(3)
				return
			}
			if i+1 < len(args) {
				str2 = args[i+1]
			}
			insert = true

		} else if len(args[i]) >= 2 && args[i][:2] == "-i" {
			if len(args[i]) == 2 || args[i][2] != '=' {
				msg(3)
				return
			}
			insertion = args[i][3:]
			if len(insertion) == 0 {
				msg(3)
				return
			}
			if i+1 < len(args) {
				str2 = args[i+1]
			}
			insert = true
		}
	}
	if order && insert {
		fmt.Println(oorder(str + insertion))
	} else if order && !insert {
		fmt.Println(oorder(str))
	} else if !order && insert {
		fmt.Println(str2 + insertion)
	} else if !order && !insert {
		if len(args) == 1 {
			fmt.Println(args[0])
		} else {
			msg(0)
			return
		}
	}
}

// quest 6

/* In Go, slicing is done with [start:end], where start
is inclusive and end is exclusive. */
