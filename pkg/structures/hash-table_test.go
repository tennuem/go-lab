package structures

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashTable(t *testing.T) {
	testData := []int{5, 4, 2, 3, 1, 9, 7, 8, 0, 6}
	s := NewHashTable()

	for _, v := range testData {
		key := strconv.Itoa(v)
		s.Set(key, v)
		assert.Equal(t, v, s.Get(key))

		s.Remove(key)
		assert.Equal(t, 0, s.Get(key))
	}
}
