package main

import "fmt"

func SwapBits(octet byte) byte {
	r := octet >> 4
	l := octet & 15
	l <<= 4
	return r | l
}

func main() {
	myByte := byte(0b10101101)
	fmt.Printf("before swap: %b\n", myByte)
	fmt.Printf("After swap : %b\n", SwapBits(myByte))
}

// checkpoint
// bits related : reversebits, maxminint, bitwise, printmemory
