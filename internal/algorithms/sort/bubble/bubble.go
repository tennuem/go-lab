package bubble

import (
	"time"

	"github.com/tennuem/go-lab/internal/algorithms/sort"
)

type Bubble struct {
	Items      []int
	Operations int
	Durations  time.Duration
}

func (b *Bubble) Result1(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		b.Durations = time.Since(start)
	}(start)

	end := len(array) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(array)-1; i++ {

			// Counting number operations
			b.Operations++

			if array[i] < array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}
		end--
	}
	return array
}

func (b *Bubble) Result(array []int) []int {

	// Time of algorithm execution
	start := time.Now()
	defer func(start time.Time) {
		b.Durations = time.Since(start)
	}(start)

	var swapped bool = true
	for swapped {
		swapped = false
		for i := 0; i < len(array)-1; i++ {

			// Counting number operations
			b.Operations++

			if array[i] < array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				swapped = true
			}
		}
	}
	return array
}

func (b *Bubble) Operation() int {
	return b.Operations
}

func (b *Bubble) Duration() time.Duration {
	return b.Durations
}

func NewBubble() sort.Sort {
	return &Bubble{}
}
