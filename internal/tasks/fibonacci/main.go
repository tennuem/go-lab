package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fib2(i))
	}
}

// fib1 by default for go
func fib1() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

// fib2 linearly recurrent relation
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}
