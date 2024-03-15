package main

import "fmt"

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(1))
	fmt.Println(IsPrime(17))
}

func IsPrime(nb int) bool {
	if nb == 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// quest 4

/* i*i<=nb ??? (finding optimization):
if n is not a prime number we will ALWAYS find i
where i<=root(n) and n%i=0

* Proof: 
Let's take n=a*b
where a, b > 1 and a, b < n and a!=b 
(a!=b remove perfect square case since they are never prime numbers)

if a < b then a < root(n) because if a > root(n) then
b > root(n) and a*b > (root(n))² = n which contradict
a*b = n.

in the same way if b > a then b < root(n).

Then if n = a*b (not a prime number) we will always
find that a or b is < root(n)
which means that there is always a (or b) where
n%a == 0 (where a < root(n))
*/
