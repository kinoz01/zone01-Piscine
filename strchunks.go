package main

import "fmt"

func printStrslice(sc []string) string {
	var f string
	for i, str := range sc {
		if i == 0 {
			f = str
			continue
		}
		f = f + " " + str
	}
	return f
}

func sliceStringEveryN2(s string, n int) []string {
	var result []string
	star := 0
	for i := 1; i < len(s); i++ {
		if i%n == 0 { // you can also use (i-star) == n
			result = append(result, s[star:i])
			star = i
		}
	}
	if star < len(s) {
		result = append(result, s[star:])
	}
	return result
}

func sliceStringEveryN(s string, n int) []string {
	// you can handle cases where n is wrong (n < 1 or n too big) but for now no need
	var result []string
	for i := 0; i < len(s); i += n {
		// Check if the remaining substring is shorter than n
		if i+n > len(s) {
			result = append(result, s[i:])
		} else {
			result = append(result, s[i:i+n])
		}
	}
	return result
}

func main() {
	s1 := "deliciousbread"
	s2 := "whay iyayyay ay"
	s3 := "heywassupdudeeverythinggood"
	fmt.Println(printStrslice(sliceStringEveryN(s1, 3)))
	fmt.Println(printStrslice(sliceStringEveryN(s2, 3)))
	fmt.Println(printStrslice(sliceStringEveryN2(s3, 4)))
}

// educative / explanation
