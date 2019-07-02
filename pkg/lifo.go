package pkg

import "sync"

type Stack interface {
	Push(v int)
	Pop() int
}

// NewStack returns a new stack
func NewStack() Stack {
	return &stack{}
}

type stack struct {
	data  []int
	count int
	mu    sync.Mutex
}

// Push adds a int to the stack
func (s *stack) Push(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, v)
	s.count++
}

// Pop returns a last int from the stack
func (s *stack) Pop() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.count == 0 {
		return 0
	}
	s.count--
	defer func() {
		s.data = s.data[:s.count]
	}()
	return s.data[s.count]
}
