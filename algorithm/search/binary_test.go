package search

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinary(t *testing.T) {
	var (
		testData     = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		searchNumber = rand.Intn(len(testData))
	)

	assert.Equal(t, Binary(testData, searchNumber), true)
}

func BenchmarkBinary(b *testing.B) {
	testData := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		Binary(testData, rand.Intn(len(testData)))
	}
}
