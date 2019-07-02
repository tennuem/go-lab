package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort1(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	res := InsertionSort1(testData)

	for i := 0; i < len(testData); i++ {
		assert.Equal(t, res[i], i)
	}
}

func TestInsertionSort2(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	res := InsertionSort2(testData)

	for i := 0; i < len(testData); i++ {
		assert.Equal(t, res[i], i)
	}
}

func BenchmarkInsertionSort1(b *testing.B) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	for i := 0; i < b.N; i++ {
		InsertionSort1(testData)
	}
}

func BenchmarkInsertionSort2(b *testing.B) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	for i := 0; i < b.N; i++ {
		InsertionSort2(testData)
	}
}
