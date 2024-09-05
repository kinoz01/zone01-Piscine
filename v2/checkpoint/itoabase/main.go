package main

func main() {

}

func ItoaBase(n, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	// Characters for digits in base systems
	digits := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D', 'E', 'F'}

	// Handle negative numbers
	negative := false
	var num uint64

	// Handle negative numbers without overflow
	if n < 0 {
		negative = true
		// Hardcode the value for -9223372036854775808 as uint64
		if n == -9223372036854775808 {
			num = 9223372036854775808 // This is uint64 and correct
		} else {
			num = uint64(-n)
		}
	} else {
		num = uint64(n)
	}

	var result string

	// Convert the number to the specified base
	for num > 0 {
		result = string(digits[num%uint64(base)]) + result
		num /= uint64(base)
	}

	// Add a minus sign for negative numbers
	if negative {
		result = "-" + result
	}

	return result
}
