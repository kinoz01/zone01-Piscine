package main

import "fmt"

func PointOne(n *int) {
	*n = 1
}

func main() {
	n := 0
	PointOne(&n)
	fmt.Println(n)
}

// Quest 3
/*i, j := 42, 2701

p := &i         // point to i
fmt.Println(*p) // read i through the pointer (42)
*p = 21         // set i through the pointer
fmt.Println(i)  // see the new value of i

p = &j         // point to j
*p = *p / 37   // divide j through the pointer
fmt.Println(j) // see the new value of j


int a=5
int *b
b=&a
*b=5
*/
