package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	hex := "0123456789abcdef"
	for i := 0; i < len(arr); i++ {
		z01.PrintRune(rune(hex[arr[i]>>4]))
		z01.PrintRune(rune(hex[arr[i]&15]))
		z01.PrintRune(' ')
		if (i-3)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		}
	}
	for _, byt := range arr {
		if byt < 32 || byt > 128 {
			z01.PrintRune('.')
		} else {
			z01.PrintRune(rune(byt))
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

// checkpoint


/************************* Explanation***************************
these two lines saved us too much code:
z01.PrintRune(rune(hex[arr[i]>>4]))
z01.PrintRune(rune(hex[arr[i]&15]))

so let's try to understand them:

The conversion from binary to hexadecimal is efficient 
and straightforward because both are base-2 systems, 
meaning hexadecimal is a base-16 system which is a power 
of 2 (2^4 = 16). This relationship allows for a direct 
and simple conversion process, unlike the conversion 
between binary and decimal, where the base-10 system does 
not align as neatly with binary's base-2 system.

Each hexadecimal digit can represent four binary digits 
(bits) because 2^4 = 16. Thus, one hexadecimal digit can 
exactly represent a range from 0000 to 1111 in binary 
(or 0 to 15 in decimal).

Since one hexadecimal digit corresponds to exactly 4 bits 
in binary, we can split a byte (8 bits) into two halves, 
with each half directly mapping to a single hexadecimal 
digit. This is why the method involves isolating the 
high and low 4-bit parts of the byte.

High-order Half (arr[i]>>4):
By shifting the byte to the right by 4 bits (>>4), you 
effectively discard the lower 4 bits and move the upper 
4 bits to the position of the lower 4 bits. This 
operation transforms the upper half of the byte into 
a value that can be directly looked up in a hexadecimal 
mapping. For any given byte, this reveals the hexadecimal 
digit for the more significant part (the "high-order" half).

Low-order Half (arr[i]&15):
Applying a bitwise AND operation with 15 (0x0F or 
00001111 in binary) masks the upper 4 bits of the 
byte, isolating the lower 4 bits. This is because 
any bit ANDed with 0 results in 0, and any bit 
ANDed with 1 remains unchanged. This operation 
reveals the hexadecimal digit for the less significant 
part (the "low-order" half).

*/
