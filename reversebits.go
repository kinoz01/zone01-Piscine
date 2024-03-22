package main

import "fmt"

func ReverseBits(oct byte) byte {
	var result byte
	for i := 1; i <= 8; i++ {
		result <<= 1
		result |= oct & 1
		oct >>= 1
	}
	return result
}

func main() {
	myByte := byte(0b10101101)
	fmt.Printf("Byte before reverse: %b\n", myByte)
	fmt.Printf("Byte after reverse : %b\n", ReverseBits(myByte))
}

// checkpoint

/********************* Exaplanation*******************/
/*
1. Shifting result to the left makes room for the new bit
by moving all bits one position to the left.
2. AND-ing oct with 1 (oct & 1) isolates the right-most bit
of oct.
3. OR-ing this bit with result adds the bit to the right-most
position of result.
4. Shifting oct to the right (oct >> 1) prepares the next
bit for processing in the next iteration.

The function essentially builds the reversed binary
number one bit at a time, shifting result to make room
for the next bit and extracting each bit from oct
starting from the right.
*/
