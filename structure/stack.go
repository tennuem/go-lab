package structure

import "sync"

//Stack https://en.wikipedia.org/wiki/Stack_(abstract_data_type)
type Stack interface {
	//Push adds an element to the collection.
	Push(v int)
	//Pop removes the most recently added element that was not yet removed.
	Pop() int
}

// NewStack returns struct implementing Stack
func NewStack() Stack {
	return &stack{}
}

type stack struct {
	data  []int
	count int
	mu    sync.Mutex
}

func (s *stack) Push(v int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, v)
	s.count++
}

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
