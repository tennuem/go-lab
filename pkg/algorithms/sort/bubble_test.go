package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort1(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	res := BubbleSort1(testData)

	for i := 0; i < len(testData); i++ {
		assert.Equal(t, res[i], i)
	}
}

func TestBubbleSort2(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	res := BubbleSort2(testData)

	for i := 0; i < len(testData); i++ {
		assert.Equal(t, res[i], i)
	}
}

func BenchmarkBubbleSort1(b *testing.B) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	for i := 0; i < b.N; i++ {
		BubbleSort1(testData)
	}
}

func BenchmarkBubbleSort2(b *testing.B) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	for i := 0; i < b.N; i++ {
		BubbleSort2(testData)
	}
}
