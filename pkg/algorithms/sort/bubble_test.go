package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	var (
		testData = []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}
		sort     = NewBubble()
	)

	res := sort.Result(testData)

	for i := 0; i < len(testData); i++ {
		assert.Equal(t, res[i], i)
	}
}
