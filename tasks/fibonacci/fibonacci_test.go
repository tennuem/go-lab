package fibonacci

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
