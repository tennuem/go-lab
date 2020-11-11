package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Fibonacci is the sum of the two preceding ones, starting from 0 and 1.
// https://en.wikipedia.org/wiki/Fibonacci_number
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func TestFibonacci(t *testing.T) {
	testData := map[int]int{
		1: 1,
		2: 1,
		3: 2,
		4: 3,
		5: 5,
	}

	for n, v := range testData {
		assert.Equal(t, Fibonacci(n), v)
	}
}
