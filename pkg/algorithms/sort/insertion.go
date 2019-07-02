package sort

import (
	"time"
)

func NewInsertion() Sort {
	return &insertionSort{}
}

type insertionSort struct {
	Operations int
	Durations  time.Duration
}

func (s *insertionSort) Result(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		s.Durations = time.Since(start)
	}(start)

	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {

			// Counting number operations
			s.Operations++

			array[j-1], array[j] = array[j], array[j-1]
		}
	}
	return array
}

func (s *insertionSort) Operation() int {
	return s.Operations
}

func (s *insertionSort) Duration() time.Duration {
	return s.Durations
}
