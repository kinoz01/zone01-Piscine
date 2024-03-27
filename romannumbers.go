package main

import "os"

// import "fmt"

func atoe(s string) int {
	var n int
	for _, digit := range s {
		n = n*10 + int(digit-'0')
	}
	return n
}

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	n := atoe(args[0])
	for _, digit := range args[0] {
		if digit > '9' || digit < '0' || n == 0 || n >= 4000 { // exiting conditions
			os.Stdout.WriteString("ERROR: cannot convert to roman digit.\n")
			return
		}
	}
	rom := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	t := []int{}
	var f string // formula breakdown string
	var r string // resulting number string
	for i := 0; i < len(rom); i++ {
		if n >= rom[i] {
			t = append(t, n/rom[i])
			n %= rom[i]
		} else {
			t = append(t, 0) // we gonna use 0 to skip unwanted roman strings index
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == 0 {
			continue
		} else {
			for j := 0; j < t[i]; j++ {
				if f != "" {
					f += "+" // add + sign to the beginning of each f term except for the begin
				}
				if len(roms[i]) == 2 {
					f += "(" + string(roms[i][1]) + "-" + string(roms[i][0]) + ")"
				} else {
					f += roms[i]
				}
				r += roms[i]
			}
		}
	}
	os.Stdout.WriteString(f + "\n")
	os.Stdout.WriteString(r + "\n")
}

// checkpoint
