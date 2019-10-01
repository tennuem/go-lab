package linked

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const linkedLinkSize = 10

func TestLinkedLinkInsertGet(t *testing.T) {
	list := NewSinglyLinkedList()
	for i := 0; i < linkedLinkSize; i++ {
		err := list.Insert(i, i*2)
		n, err := list.Get(i)
		assert.NoError(t, err)
		assert.Equal(t, i*2, n.Value)
	}
}

func TestLinkedListRemoveGet(t *testing.T) {
	list := NewSinglyLinkedList()
	for i := 0; i < linkedLinkSize; i++ {
		err := list.Insert(i, i*2)
		err = list.Remove(i)
		_, err = list.Get(i)
		assert.EqualError(t, err, ErrPositionNotFound.Error())
	}
}
