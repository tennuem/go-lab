package sort

import (
	"time"
)

func NewBubble() Sort {
	return &bubbleSort{}
}

type bubbleSort struct {
	Items      []int
	Operations int
	Durations  time.Duration
}

func (s *bubbleSort) Result1(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		s.Durations = time.Since(start)
	}(start)

	end := len(array) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(array)-1; i++ {

			// Counting number operations
			s.Operations++

			if array[i] < array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		end--
	}
	return array
}

func (s *bubbleSort) Result(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		s.Durations = time.Since(start)
	}(start)

	var swapped bool = true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {

			// Counting number operations
			s.Operations++

			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}
	return array
}

func (s *bubbleSort) Operation() int {
	return s.Operations
}

func (s *bubbleSort) Duration() time.Duration {
	return s.Durations
}
