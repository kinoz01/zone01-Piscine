package main

func ForEach(f func(int), a []int) {
	for _, n := range a {
		f(n)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}

// quest 9
