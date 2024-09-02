package main

import (
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 || len(os.Args[1]) < 2 || len(os.Args[2]) == 0 {
		return
	}

	if os.Args[1][0] != '(' || os.Args[1][len([]rune(os.Args[1]))-1] != ')' {
		return
	}

	s := Split(os.Args[1][1:len([]rune(os.Args[1]))-1], '|', false)
	if s == nil {
		return
	}
	f := Split(os.Args[2], ' ', true)

	result := []string{}

	for _, word := range f {
		for _, reg := range s {
			if Contains(word, reg) {
				if IsNotLetter(rune(word[len(word)-1])) {
					result = append(result, word[0:len(word)-1])
				} else {
					result = append(result, word)
				}
			}
		}
	}
	for i, w := range result {
		os.Stdout.WriteString(Itoi(i+1) + ": " + w + "\n")
	}
}

func Split(s string, sep rune, field bool) (res []string) {

	if s == "" {
		return nil
	}
	var temp []rune
	for _, r := range s {
		if r == sep {
			res = append(res, string(temp))
			temp = nil
			continue
		}
		temp = append(temp, r)
	}
	res = append(res, string(temp))

	if field {
		temp := res
		res = nil
		for _, str := range temp {
			if str != "" {
				res = append(res, str)
			}
		}
		return res
	}
	return res
}

func IsNotLetter(r rune) bool {
	return !((r <= 'z' && r >= 'a') || (r <= 'Z' && r >= 'A'))
}

func Itoi(n int) (s string) {
	for n > 0 {
		s = string(rune(n%10)+'0') + s
		n /= 10
	}
	return s
}

func Contains(s string, reg string) bool {
	for i := 0; i <= len(s)-len(reg); i++ {
		if s[i:i+len(reg)] == reg {
			return true
		}
	}
	return false
}
