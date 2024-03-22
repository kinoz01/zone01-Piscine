package main

import (
	"fmt"
)

func main() {

	var x = 9
	var y = 8

	fmt.Printf("x = %b\n", x)
	fmt.Printf("y = %b\n", y)
	fmt.Printf("x & y is %b\n", x&y)
	fmt.Printf("x | y is %b\n", x|y)
	fmt.Printf("x ^ y is %b\n", x^y)
	fmt.Printf("x << 2 is %b\n", x<<2)
	fmt.Printf("x >> 2 is %b\n", x>>2)
	/* Note: The result of NOT operation depends on the
	system's word size (e.g., 32-bit, 64-bit)
	Here, we mask the result to the lower 8 bits for a
	clearer demonstration */
	notResult := ^11 & 0xFF // ~1011
	fmt.Printf("NOT Operation: ~11 = ~1011 = %b\n", notResult)
}

/********************* Explanation *******************/
/*
Bitwise operators perform operations on binary
representations of numbers at the bit level. Here's an
explanation of each bitwise operator with examples:

1. AND (&)
Operation: Sets each bit to 1 if both bits are 1.
0 & 0 = 0
0 & 1 = 0
1 & 0 = 0
1 & 1 = 1
Example: 1011 (11 in decimal) & 1100 (12 in decimal)
= 1000 (8 in decimal)
1011
1100
----
1000

2. OR (|)
Operation: Sets each bit to 1 if at least one of the
two bits is 1.
0 | 0 = 0
0 | 1 = 1
1 | 0 = 1
1 | 1 = 1
Example: 1011 (11 in decimal) | 1100 (12 in decimal)
= 1111 (15 in decimal)
1011
1100
----
1111

3. XOR (Exclusive OR) (^)
Operation: Sets each bit to 1 if only one of the two
bits is 1.
0 ^ 0 = 0
0 ^ 1 = 1
1 ^ 0 = 1
1 ^ 1 = 0
Example: 1011 (11 in decimal) ^ 1100 (12 in decimal)
= 0111 (7 in decimal)
1011
1100
----
0111

4. NOT (~)
Operation: Inverts all the bits (turns 0 into 1 and
1 into 0).
~0 = 1
~1 = 0
Example: ~1011 (11 in decimal) = 0100 (in 4-bit
representation, this would actually be the complement,
so in full 8-bit it's more like ~00001011 = 11110100;
the result depends on the number of bits your system
uses for representations)
~1011
-----
0100 (This example is simplified; actual
result depends on bit width)

5. Left Shift (<<)
Operation: Shifts the bits to the left, padding with
zeros on the right. This is equivalent to multiplying by
2^(n) (where n is the number of shifts).
Example: 1011 (11 in decimal) << 2 = 101100 (44 in decimal)
1011 << 2
---------
101100

6. Right Shift (>>)
Operation: Shifts the bits to the right. For signed
numbers, the behavior can depend on the system
(logical vs. arithmetic shift). A logical shift fills
the leftmost bits with zeros, while an arithmetic shift
fills with the sign bit (maintaining the sign of
the number). For unsigned numbers, it's just a
zero-fill. This is equivalent to dividing by 2^(n)
and discarding the remainder.
Example (assuming logical shift): 1011 (
11 in decimal) >> 2 = 0010 (2 in decimal)
1011 >> 2
---------
0010
*/

/**********Two's complement representation *************/
/* Two's complement representation is a method for
encoding signed binary numbers in a way that makes
addition, subtraction, and comparison operations
simple and consistent, especially for computer hardware.
It's the most common method for representing signed
integers in computers.

* To Find the Two's Complement of a Positive Number:
For positive numbers and zero, the two's complement
representation is the same as the regular binary
representation. The leftmost bit (most significant bit,
MSB) is 0, indicating that the number is positive.

* To Find the Two's Complement of a Negative Number:
1. Start with the absolute value of the number in binary
form.
2. Invert the bits — change all 0s to 1s and all 1s to 0s.
This is known as the one's complement.
3. Add 1 to the result of step 2.

Example: Converting -6 to Two's Complement
(8-bit Representation)
1. Absolute value in binary: 00000110 (6 in binary)
2. Invert the bits: 11111001
3. Add 1: 11111001 + 1 = 11111010
Thus, -6 in two's complement (8-bit) is 11111010.

Example: Converting 11111010 Back to Decimal
1. Recognize it's a negative number (MSB is 1).
2. Convert back to positive by reversing the two's
complement process:
Subtract 1: 11111001
Invert the bits: 00000110 (which is 6 in binary)
3. The original number is -6.
*/

/*********** Explanation of & 0xFF in the code **********/
/*In Go, the ^ operator used with a single operand is
the bitwise complement operator, which inverts all bits
of its operand. When you apply it to an integer like 11
(which is 00001011 in 8-bit binary), the bitwise NOT
operation ^11 inverts all bits, resulting in a binary
number where all bits are the inverse of 11. This
operation is performed on the full width of the type of
the number. If 11 is considered an int type, and assuming
int is 32 bits on your system, the binary representation
of 11 would actually be padded with zeros to fit this
width: 0000...0001011. After the ^ operation, all bits
are inverted, giving you 1111...1110100.

Removing 0xFF from the operation changes the context
significantly:

With 0xFF masking: The operation ^11 & 0xFF first inverts
all bits of 11, then performs a bitwise AND with 0xFF
(11111111 in binary). This masks the result to the last
8 bits, effectively showing you the "positive"
representation of the inverted value in an 8-bit context,
which would be within the range 0 to 255.
Without 0xFF masking: You see the result of the bitwise
NOT operation on the entire integer width (^11), which,
for a 32-bit integer, involves inverting all 32 bits.
Since the highest bit (sign bit) becomes 1, the number is
interpreted as a negative number in two's complement
representation.

The value -1100 you see is the decimal representation
of the binary result of the NOT operation, interpreted
as a negative number because the most significant bit
(which indicates the sign in a two's complement system)
is set to 1. This is the two's complement of the binary
result after inverting all bits of 11.
*/
