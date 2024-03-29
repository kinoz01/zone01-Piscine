package main

import "os"

// import "fmt"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	if args[0] == "" {
		os.Stdout.WriteString("\n")
		return
	}
	strSlice := splitt(args[0], " ")
	result := strSlice[len(strSlice)-1]
	for i := len(strSlice) - 2; i >= 0; i-- {
		result = result + " " + strSlice[i]
	}
	os.Stdout.WriteString(result + "\n")
}

func splitt(s, sp string) []string {
	strs := []string{}
	index := 0
	for i := 0; i+len(sp) <= len(s); i++ {
		if s[i:i+len(sp)] == sp {
			if s[index:i] != "" {
				strs = append(strs, s[index:i])
			}
			index = i + len(sp)
		} else if i+len(sp) == len(s) { // if len s[i:i+len(sp)]!= sep and i+len(sp)=len(s) (WHIICH MEANS we are at the last iteration) WHIICH MEANS if len s[len(s)-len(sep):len(s))]!= sep WHIICH MEANS if the end of the string doesn't end with a sep
			strs = append(strs, s[index:]) // in this case append all to the end
		}
	}
	// fmt.Printf("%#v\n", strs) // this line allow us to see clearly how and why we added the condition "if s[index:i]!=""{" above
	return strs
}

// checkpoint
