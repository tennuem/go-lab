package structures

import "sync"

// List https://en.wikipedia.org/wiki/Dynamic_array
type List interface {
	Push(v int)
	Pop() int
	UnShift(v int)
	Shift() int
}

// NewList returns new struct implementing List
func NewList() List {
	return &list{}
}

type list struct {
	mu     sync.Mutex
	length int
	data   []int
}

func (l *list) Push(v int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = append(l.data, v)
	l.length++
}

func (l *list) Pop() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.length == 0 {
		return 0
	}
	i := l.length - 1
	v := l.data[i]
	l.data = l.data[:i]
	l.length--
	return v
}

func (l *list) UnShift(v int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.data = append([]int{v}, l.data...)
	l.length++
}

func (l *list) Shift() int {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.length == 0 {
		return 0
	}
	v := l.data[0]
	l.data = l.data[1:]
	l.length--
	return v
}
