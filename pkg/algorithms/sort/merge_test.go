package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	res := MergeSort(testData)

	for i := 0; i < len(testData); i++ {
		assert.Equal(t, res[i], i)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}

	for i := 0; i < b.N; i++ {
		MergeSort(testData)
	}
}
