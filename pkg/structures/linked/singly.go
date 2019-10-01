package linked

import (
	"github.com/pkg/errors"
)

var (
	ErrPositionNotFound = errors.New("position >= length")
	ErrHeadNotFound     = errors.New("head not found")
)

type Node struct {
	Value int
	Next  *Node
}

type SinglyLinkedList struct {
	len  int
	root *Node
}

// NewSinglyLinkedList https://en.wikipedia.org/wiki/Linked_list#Singly_linked_list
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (l *SinglyLinkedList) Insert(position int, value int) error {
	if position > 0 && l.len == 0 {
		return ErrHeadNotFound
	}

	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	if position == 0 {
		newNode.Next = l.root
		l.root = newNode
		l.len++
		return nil
	}

	prev, err := l.Get(position - 1)
	if err != nil {
		return err
	}
	current := prev.Next

	newNode.Next = current
	prev.Next = newNode

	l.len++

	return nil
}

func (l *SinglyLinkedList) Get(position int) (*Node, error) {
	if position >= l.len {
		return nil, ErrPositionNotFound
	}
	current := l.root
	for i := 0; i < position; i++ {
		current = current.Next
	}
	return current, nil
}

func (l *SinglyLinkedList) Remove(position int) error {
	if position == 0 {
		l.root = l.root.Next
		l.len--
		return nil
	}

	prev, err := l.Get(position - 1)
	if err != nil {
		return err
	}
	prev.Next = prev.Next.Next
	l.len--

	return nil
}
