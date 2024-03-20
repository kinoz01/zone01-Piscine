package main

import "github.com/01-edu/z01"

func main() {
	PrintNbr2(-123)
	PrintNbr2(0)
	PrintNbr2(123)
	z01.PrintRune('\n')
}

func PrintNbr2(n int) {
	var nbr int
	var r string
	var f string
	if n == 0 {
		z01.PrintRune('0') // if n is equal 0 print '0' and return
		return
	}
	if n < 0 {
		nbr = -n // use nbr positive and keep handling with sign to later
	} else {
		nbr = n
	}
	for nbr > 0 {
		r = string(nbr%10+'0') + r
		nbr /= 10
	}
	if n < 0 {
		f = "-" + r // we gonna use to print minus sign
		for _, char := range f {
			z01.PrintRune(char)
		}
	} else {
		for _, char := range r {
			z01.PrintRune(char)
		}
	}
}

/****************** NEW Favorite way**************/
// func printnum(n int) {
// 	if n<0 {
// 		z01.PrintRune('-')
// 		n=-n
// 	}
// 	if n/10!=0{
// 		printnum(n/10)
// 	}
// 	z01.PrintRune(rune(n%10)+'0')
// }

// quest 2 better way for printnbr (itoa) you can make

/****************EXPLANATION**************/
/*The character code of '0' is 48 in ASCII.
When you add the integer result of nbr % 10 to the
character code of '0' (48), you get the character code
of the resulting digit.*/

/*In Go, string(65) converts the integer 65 into its
corresponding character based on the Unicode code point.
The integer 65 represents the letter "A" in the ASCII
and Unicode character sets. Therefore, string(65) in
Go would result in the string "A".*/

// https://theasciicode.com.ar

/* the same prinsiple will apply if you use rune(65).
However While string(65) results in a string containing
the character "A", rune(65) results in a rune, which is
a numeric type representing the character "A"'s code
point in Unicode. If you print it directly, Go will
display the numeric value of the rune, unless you convert
it to a string or format it to display as a character.*/
