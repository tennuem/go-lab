package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Factorial is the product of all positive integers less than or equal to n.
// https://en.wikipedia.org/wiki/Factorial
func Factorial(n int) int {
	if n > 1 {
		return n * Factorial(n-1)
	}
	return 1
}

func TestFactorial(t *testing.T) {
	testData := map[int]int{
		0: 1,
		1: 1,
		2: 2,
		3: 6,
		4: 24,
		5: 120,
	}

	for n, v := range testData {
		assert.Equal(t, Factorial(n), v)
	}
}
