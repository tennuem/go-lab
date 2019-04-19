package main

import (
	"fmt"
	"sync"
	"time"
)

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
	return s.data[s.count]
}

func main() {
	var wg sync.WaitGroup
	s := NewStack()

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			s.Push(i)
		}
	}()

	// second goroutine read from stack, bfore first goroutine write to the stack
	time.Sleep(time.Second * 1)

	go func() {
		for {
			v := s.Pop()
			if v == 0 {
				wg.Done()
				break
			}
			fmt.Println(v)
		}
	}()

	wg.Wait()
}
