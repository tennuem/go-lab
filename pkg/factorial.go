package pkg

// Factorial is the product of all positive integers less than or equal to n.
// https://en.wikipedia.org/wiki/Factorial
func Factorial(n int) int {
	if n > 1 {
		return n * Factorial(n-1)
	}
	return 1
}
