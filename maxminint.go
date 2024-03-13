package main

import "fmt"

func main() {
	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1
	fmt.Println(maxInt)
	fmt.Println(minInt)
}
/*************************************
In a 64-bit system:

* Bits: You have 64 bits to represent an integer.
* Sign Bit: The first bit is the sign bit (0 for positive, 1 
for negative).
* Magnitude Bits: The remaining 63 bits represent the magnitude of the number.

//* Maximum Positive Integer (2^63 - 1):

* Binary Representation: 01111111 11111111 11111111 11111111 
11111111 11111111 11111111 11111111
* Explanation: The leftmost bit (sign bit) is 0, 
indicating the number is positive. All the other bits are 
set to 1, giving the largest possible magnitude for a 
positive number with 63 bits.
*Decimal Value: 9223372036854775807

//* Minimum Negative Integer (-2^63):

Binary Representation: 10000000 00000000 00000000 00000000 
00000000 00000000 00000000 00000000
* Explanation: The leftmost bit (sign bit) is 1, 
indicating the number is negative. All the other bits are 
set to 0. This is the binary representation of -2^63 in 
two's complement, which is the way negative numbers are 
represented.
* Decimal Value: -9223372036854775808

//* Why the Difference?

The two's complement system is designed to make arithmetic 
operations simpler and to have a unique representation for 
zero. In two's complement:

Positive Numbers: Are represented as themselves.
Negative Numbers: Are represented by inverting the bits 
of their positive counterparts and adding 1.
This representation allows for an additional negative 
number because the bit pattern that represents -2^63 does 
not have an equivalent positive number. The highest 
positive number uses all 63 magnitude bits set to 1, but 
adding 1 to this number to try and get a positive 
representation of 2^63 would require a 64th magnitude bit, 
which is not available because the 64th bit is reserved 
for the sign.

Thus, in a 64-bit system with one bit reserved for the sign, you can 
represent numbers from -2^63 to 2^63 - 1, leading to the 
observed asymmetry between the maximum positive 
and minimum negative values.
****************************************/


/*******************************************
^uint(0):

uint is an unsigned integer type, meaning it can only 
represent non-negative numbers. The size (in bits) of 
uint can vary depending on the system, but it's commonly 
32 or 64 bits.
0 is just the number zero, represented in uint type.
^ is the bitwise NOT operator. Applying it to 0 inverts 
all bits of 0, turning them into 1. So, if uint is 32 
bits, ^uint(0) results in 0xFFFFFFFF (all bits set to 1).
If uint is 64 bits, it results in 0xFFFFFFFFFFFFFFFF.
^uint(0) >> 1:

>> is the bitwise right shift operator. Shifting ^uint(0) 
(which is all bits set to 1) to the right by 1 position 
effectively divides the number by 2 and also discards the 
least significant bit. This operation changes the leftmost 
bit from 1 to 0, making sure the result is the largest 
possible number that can be represented by an int type 
(which is a signed integer type).
For a 32-bit system, this shifts from 0xFFFFFFFF to
0x7FFFFFFF.
For a 64-bit system, this shifts from 0xFFFFFFFFFFFFFFFF 
to 0x7FFFFFFFFFFFFFFF.

int(^uint(0)>>1):

This final part casts the result to an int type. Because 
we've shifted right by 1, the highest bit (which would
indicate a negative number in a signed integer) is 
now 0, and thus, the number is the largest positive 
integer that can be stored in an int.
The result is the maximum positive value that an int can 
hold on the current system. This technique is 
system-agnostic, meaning it automatically adjusts to the 
size of int on the system where the code is compiled, 
whether it's 32-bit or 64-bit.
*****************************************************/
