package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}
	s := NewQueue()

	for _, v := range testData {
		s.Enqueue(v)
		assert.Equal(t, v, s.Dequeue())
	}
}
