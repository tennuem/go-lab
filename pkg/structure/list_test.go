package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}
	l := NewList()

	for _, v := range testData {
		l.Push(v)
		assert.Equal(t, v, l.Pop())

		l.UnShift(v)
		assert.Equal(t, v, l.Shift())
	}
}
