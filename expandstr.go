package main

import "os"

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	arr := []rune{}
	runeA := []rune(args[0])
	for i := 0; i < len(runeA); i++ {
		if runeA[i] != ' ' && runeA[i] != '\t' {
			arr = append(arr, runeA[i]) // append all characters except whitespaces
		}
		if (runeA[i] != ' ' && runeA[i] != '\t') && i+1 < len(runeA) && (runeA[i+1] == ' ' || runeA[i+1] == '\t') {
			arr = append(arr, ',') // Append a comma only if the current rune is not a space or a tab, and the next rune is a space or a tab (i used comma we can use any other mark)
		}
	}
	r := []rune{}
	if arr[len(arr)-1] == ',' {
		r = arr[0 : len(arr)-1] // remove the last comma appended in case the string ends whith whitespace
	} else {
		r = arr
	}
	strs := []string{}
	ind := 0 // appending from ind to where we found ',' (comma / our mark)
	for i := 0; i < len(r); i++ {
		if r[i] == ',' {
			strs = append(strs, string(r[ind:i]))
			ind = i + 1
		}
	}
	strs = append(strs, string(r[ind:])) // append the last part of the runes
	// fmt.Println(strs)
	final := strs[0]
	for i := 1; i < len(strs); i++ {
		final = final + "   " + strs[i] // we can use any other sperator instead of "   "
	}
	os.Stdout.WriteString(final + "\n")
}

// checkpoint / the lines from 11 to 25 can be replaced by a simple codition like we did in split (fik fik t appendi men index l i men be3d)
