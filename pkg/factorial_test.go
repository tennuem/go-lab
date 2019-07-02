package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
