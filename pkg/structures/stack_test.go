package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	s := NewStack()

	for i := 0; i < 10; i++ {
		s.Push(i)
		assert.Equal(t, s.Pop(), i)
	}
}
