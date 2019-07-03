package pkg

// Fibonacci is the sum of the two preceding ones, starting from 0 and 1.
// https://en.wikipedia.org/wiki/Fibonacci_number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
