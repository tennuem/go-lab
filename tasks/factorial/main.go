package main

import "fmt"

func main() {
	fmt.Println(factorialCompute(10))
}

func factorialCompute(n int) int {
	if n > 1 {
		return n * factorialCompute(n-1)
	}
	return 1
}
