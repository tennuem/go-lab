package structures

import "sync"

type Queue interface {
	Enqueue(value int)
	Dequeue() int
	Peek() int
}

func NewQueue() Queue {
	return &queue{}
}

type queue struct {
	mu     sync.Mutex
	length int
	data   []int
}

func (s *queue) Enqueue(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, value)
	s.length++
}

func (s *queue) Dequeue() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.length == 0 {
		return 0
	}
	v := s.data[0]
	s.data = s.data[1:]
	s.length--
	return v
}

func (s *queue) Peek() int {
	return s.data[0]
}
