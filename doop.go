package main

import "os"

// import "fmt"

const (
	maxInt = int(^uint(0) >> 1)
	minInt = -maxInt - 1
)

func main() {
	// max int in 64bit:
	// 9223372036854775807 (len(s)=19)
	// fmt.Println(atoii("92233720368547758099")) atoii here won't refuse overflow number which forced me to hardcode a check for overflow in the string
	// fmt.Println(itoaa(9223372036854775808))
	// fmt.Println(maxInt-1)
	// fmt.Println(minInt+1)
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	result := doop(arg[0], arg[1], arg[2])
	os.Stdout.WriteString(result) // we use standard output with writestring instead of fmt
}

func doop(s0, o, s2 string) string {
	var modulo int
	var div int
	var plus int
	var moin int
	var product int
	first := true
	first2 := true

	if o != "+" && o != "-" && o != "/" && o != "*" && o != "%" {
		return ""
	}
	for i := 0; i < len(s0); i++ {
		if (s0[0] == '+' || s0[0] == '-') && first {
			first = false
			continue
		}
		if s0[i] > '9' || s0[i] < '0' {
			return ""
		}
	}
	for i := 0; i < len(s2); i++ {
		if (s2[0] == '+' || s2[0] == '-') && first2 {
			first2 = false
			continue
		}
		if s2[i] > '9' || s2[i] < '0' {
			return ""
		}
	}
	/********** in this lines i need to remove overflow before  using atoii cuz the number will change ***********/
	if len(s0) > 19 || len(s2) > 19 {
		return ""
	}
	if s0[0] == '-' || s0[0] == '+' {
		if s0[1] >= '9' && s0[2] >= '2' && s0[3] >= '2' && s0[4] >= '3' && s0[5] >= '3' && s0[6] >= '7' && s0[7] >= '2' && s0[8] >= '0' && s0[9] >= '3' && s0[10] >= '6' && s0[11] >= '8' && s0[12] >= '5' && s0[13] >= '4' && s0[14] >= '7' && s0[15] >= '7' && s0[16] >= '5' && s0[17] >= '8' && s0[18] >= '0' && s0[19] >= '7' {
			return ""
		}
	} else {
		if s0[0] >= '9' && s0[1] >= '2' && s0[2] >= '2' && s0[3] >= '3' && s0[4] >= '3' && s0[5] >= '7' && s0[6] >= '2' && s0[7] >= '0' && s0[8] >= '3' && s0[9] >= '6' && s0[10] >= '8' && s0[11] >= '5' && s0[12] >= '4' && s0[13] >= '7' && s0[14] >= '7' && s0[15] >= '5' && s0[16] >= '8' && s0[17] >= '0' && s0[18] >= '7' {
			return ""
		}
	}
	if s2[0] == '-' || s2[0] == '+' {
		if s2[1] >= '9' && s2[2] >= '2' && s2[3] >= '2' && s2[4] >= '3' && s2[5] >= '3' && s2[6] >= '7' && s2[7] >= '2' && s2[8] >= '0' && s2[9] >= '3' && s2[10] >= '6' && s2[11] >= '8' && s2[12] >= '5' && s2[13] >= '4' && s2[14] >= '7' && s2[15] >= '7' && s2[16] >= '5' && s2[17] >= '8' && s2[18] >= '0' && s2[19] >= '7' {
			return ""
		}
	} else {
		if s2[0] >= '9' && s2[1] >= '2' && s2[2] >= '2' && s2[3] >= '3' && s2[4] >= '3' && s2[5] >= '7' && s2[6] >= '2' && s2[7] >= '0' && s2[8] >= '3' && s2[9] >= '6' && s2[10] >= '8' && s2[11] >= '5' && s2[12] >= '4' && s2[13] >= '7' && s2[14] >= '7' && s2[15] >= '5' && s2[16] >= '8' && s2[17] >= '0' && s2[18] >= '7' {
			return ""
		}
	}
	/**************************************************/

	n1 := atoii(s0)
	n2 := atoii(s2)
	if o == "%" {
		if n2 == 0 {
			return "No modulo by 0\n"
		}
		modulo = n1 % n2
		return itoaa(modulo) + "\n"
	}
	if o == "/" {
		if n2 == 0 {
			return "No division by 0\n"
		}
		div = n1 / n2
		return itoaa(div) + "\n"
	}
	if o == "+" {
		if addCheck(n1, n2) {
			return ""
		}
		plus = n1 + n2
		return itoaa(plus) + "\n"
	}
	if o == "*" {
		if multipCheck(n1, n2) {
			return ""
		}
		if n2 == 0 || n1 == 0 {
			return "0\n"
		}
		product = n1 * n2
		return itoaa(product) + "\n"
	}
	if o == "-" {
		if subCheck(n1, n2) {
			return ""
		}
		moin = n1 - n2
		return itoaa(moin) + "\n"
	}
	return ""
}

func itoaa(nbr int) string {
	var r string
	var n int
	if nbr == 0 {
		return "0"
	}
	if nbr < 0 {
		n = -nbr
	} else {
		n = nbr
	}
	for n > 0 {
		r = string(n%10+'0') + r
		n /= 10
	}
	if nbr < 0 {
		return "-" + r
	}
	return r
}

func atoii(s string) int {
	var n int
	sign := 1
	for _, digit := range s {
		if digit == '+' {
			continue
		}
		if digit == '-' {
			sign = -1
			continue
		}
		n = n*10 + int(digit-'0')
	}
	return n * sign
}

/************** checking if operation will give overflow ************************/
func addCheck(a, b int) bool {
	if a > 0 && b > 0 && a > maxInt-b {
		return true // overflow
	}
	if a < 0 && b < 0 && a < minInt+b {
		return true // underflow
	}
	return false // no overflow
}

func subCheck(a, b int) bool {
	if a < 0 && b > 0 && minInt+a < -b {
		return true
	}
	if a > 0 && b < 0 && maxInt-a < -b {
		return true
	}
	return false
}

func multipCheck(a, b int) bool {
	if a > 0 && b > 0 && a*b > maxInt {
		return true
	}
	if a < 0 && b < 0 && a*b < maxInt {
		return true
	}
	if a < 0 && b > 0 && a*b < minInt {
		return true
	}
	if a > 0 && b < 0 && b*a < minInt {
		return true
	}
	return false
}

// quest 9 / checkpoint
